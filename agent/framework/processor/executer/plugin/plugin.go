// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package plugin contains general interfaces and types relevant to plugins.
// It also provides the methods for registering plugins.
package plugin

import (
	"runtime/debug"
	"sync"

	"github.com/aws/amazon-ssm-agent/agent/appconfig"
	"github.com/aws/amazon-ssm-agent/agent/context"
	"github.com/aws/amazon-ssm-agent/agent/framework/runpluginutil"
	"github.com/aws/amazon-ssm-agent/agent/plugins/configurecontainers"
	"github.com/aws/amazon-ssm-agent/agent/plugins/configurepackage"
	"github.com/aws/amazon-ssm-agent/agent/plugins/dockercontainer"
	"github.com/aws/amazon-ssm-agent/agent/plugins/downloadcontent"
	"github.com/aws/amazon-ssm-agent/agent/plugins/inventory"
	"github.com/aws/amazon-ssm-agent/agent/plugins/lrpminvoker"
	"github.com/aws/amazon-ssm-agent/agent/plugins/refreshassociation"
	"github.com/aws/amazon-ssm-agent/agent/plugins/rundocument"
	"github.com/aws/amazon-ssm-agent/agent/plugins/runscript"
	"github.com/aws/amazon-ssm-agent/agent/plugins/updatessmagent"
	"github.com/aws/amazon-ssm-agent/agent/session/plugins/interactivecommands"
	"github.com/aws/amazon-ssm-agent/agent/session/plugins/noninteractivecommands"
	"github.com/aws/amazon-ssm-agent/agent/session/plugins/port"
	"github.com/aws/amazon-ssm-agent/agent/session/plugins/sessionplugin"
	"github.com/aws/amazon-ssm-agent/agent/session/plugins/standardstream"
)

var once sync.Once

// registeredPlugins stores the registered plugins.
var registeredPlugins *runpluginutil.PluginRegistry

type CloudWatchFactory struct {
}

func (f CloudWatchFactory) Create(context context.T) (runpluginutil.T, error) {
	return lrpminvoker.NewPlugin(context, appconfig.PluginNameCloudWatch)
}

type InventoryGathererFactory struct {
}

func (f InventoryGathererFactory) Create(context context.T) (runpluginutil.T, error) {
	return inventory.NewPlugin(context)
}

type RunPowerShellFactory struct {
}

func (f RunPowerShellFactory) Create(context context.T) (runpluginutil.T, error) {
	return runscript.NewRunPowerShellPlugin(context)
}

type UpdateAgentFactory struct {
}

func (f UpdateAgentFactory) Create(context context.T) (runpluginutil.T, error) {
	return updatessmagent.NewPlugin(context)
}

type ConfigureContainerFactory struct {
}

func (f ConfigureContainerFactory) Create(context context.T) (runpluginutil.T, error) {
	return configurecontainers.NewPlugin(context)
}

type RunDockerFactory struct {
}

func (f RunDockerFactory) Create(context context.T) (runpluginutil.T, error) {
	return dockercontainer.NewPlugin(context)
}

type ConfigurePackageFactory struct {
}

func (f ConfigurePackageFactory) Create(context context.T) (runpluginutil.T, error) {
	return configurepackage.NewPlugin(context)
}

type RefreshAssociationFactory struct {
}

func (f RefreshAssociationFactory) Create(context context.T) (runpluginutil.T, error) {
	return refreshassociation.NewPlugin(context)
}

type DownloadContentFactory struct {
}

func (d DownloadContentFactory) Create(context context.T) (runpluginutil.T, error) {
	return downloadcontent.NewPlugin(context)
}

type RunDocumentFactory struct {
}

func (r RunDocumentFactory) Create(context context.T) (runpluginutil.T, error) {
	return rundocument.NewPlugin(context)
}

type SessionPluginFactory struct {
	newPluginFunc sessionplugin.NewPluginFunc
}

func (f SessionPluginFactory) Create(context context.T) (runpluginutil.T, error) {
	return sessionplugin.NewPlugin(context, f.newPluginFunc)
}

// RegisteredWorkerPlugins returns all registered core modules.
func RegisteredWorkerPlugins(context context.T) runpluginutil.PluginRegistry {

	defer func() {
		if msg := recover(); msg != nil {
			context.Log().Errorf("Agent failed while getting registered worker plugins %v!", msg)
			context.Log().Errorf("Stacktrace:\n%s", debug.Stack())
		}
	}()

	once.Do(func() {
		loadWorkers(context)
	})
	return *registeredPlugins
}

// RegisteredSessionWorkerPlugins returns all registered session plugins.
func RegisteredSessionWorkerPlugins(context context.T) runpluginutil.PluginRegistry {
	once.Do(func() {
		loadSessionPlugins(context)
	})
	return *registeredPlugins
}

// loadWorkers loads all worker plugins that are invokers for interacting with long running plugins and
// then all standard worker plugins (if there are any conflicting names, the standard worker plugin wins)
func loadWorkers(context context.T) {
	plugins := runpluginutil.PluginRegistry{}

	//Long running plugins are handled by lrpm. lrpminvoker is a worker plugin that can communicate with lrpm.
	//that's why all long running plugins are first handled by lrpminvoker - which then hands off the work to lrpm.
	plugins[appconfig.PluginNameCloudWatch] = CloudWatchFactory{}

	for key, value := range loadPlatformIndependentPlugins(context) {
		plugins[key] = value
		context.Log().Infof("Successfully loaded platform independent plugin %v", key)
	}

	for key, value := range loadPlatformDependentPlugins(context) {
		plugins[key] = value
		context.Log().Infof("Successfully loaded platform dependent plugin %v", key)
	}

	registeredPlugins = &plugins
}

// loadSessionPlugins loads all session plugins
func loadSessionPlugins(context context.T) {
	var sessionPlugins = runpluginutil.PluginRegistry{}
	var knownSessionPlugins = map[string]runpluginutil.PluginFactory{
		appconfig.PluginNameStandardStream:         SessionPluginFactory{standardstream.NewPlugin},
		appconfig.PluginNameInteractiveCommands:    SessionPluginFactory{interactivecommands.NewPlugin},
		appconfig.PluginNamePort:                   SessionPluginFactory{port.NewPlugin},
		appconfig.PluginNameNonInteractiveCommands: SessionPluginFactory{noninteractivecommands.NewPlugin},
	}
	if len(context.AppConfig().EnablePlugins) != 0 {
		for _, enabledPlugin := range context.AppConfig().EnablePlugins {
			if _, ok := knownPlatformDependentPlugins[enabledPlugin]; ok {
				sessionPlugins[enabledPlugin] = knownPlatformDependentPlugins[enabledPlugin]
			}
		}
	} else {
		sessionPlugins = knownSessionPlugins
	}

	registeredPlugins = &sessionPlugins
}

// loadPlatformIndependentPlugins registers plugins common to all platforms
func loadPlatformIndependentPlugins(context context.T) runpluginutil.PluginRegistry {
	var workerPlugins = runpluginutil.PluginRegistry{}
	var knownPlatformDependentPlugins = map[string]runpluginutil.PluginFactory{
		inventory.Name(): InventoryGathererFactory{},
		appconfig.PluginNameAwsRunPowerShellScript: RunPowerShellFactory{},
		updatessmagent.Name():                      UpdateAgentFactory{},
		configurecontainers.Name():                 ConfigureContainerFactory{},
		dockercontainer.Name():                     RunDockerFactory{},
		refreshassociation.Name():                  RefreshAssociationFactory{},
		configurepackage.Name():                    ConfigurePackageFactory{},
		downloadcontent.Name():                     DownloadContentFactory{},
		rundocument.Name():                         RunDocumentFactory{},
	}
	if len(context.AppConfig().EnablePlugins) != 0 {
		for _, enabledPlugin := range context.AppConfig().EnablePlugins {
			if _, ok := knownPlatformDependentPlugins[enabledPlugin]; ok {
				workerPlugins[enabledPlugin] = knownPlatformDependentPlugins[enabledPlugin]
			}
		}
	} else {
		workerPlugins = knownPlatformDependentPlugins
	}

	return workerPlugins
}
