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
//
//go:build darwin || freebsd || linux || netbsd || openbsd
// +build darwin freebsd linux netbsd openbsd

package plugin

import (
	"strings"

	"github.com/aws/amazon-ssm-agent/agent/appconfig"
	"github.com/aws/amazon-ssm-agent/agent/context"
	"github.com/aws/amazon-ssm-agent/agent/framework/runpluginutil"
	"github.com/aws/amazon-ssm-agent/agent/plugins/domainjoin"
	"github.com/aws/amazon-ssm-agent/agent/plugins/endpoint"
	"github.com/aws/amazon-ssm-agent/agent/plugins/runscript"
)

var knownPlatformDependentPlugins = map[string]runpluginutil.PluginFactory{
	appconfig.PluginNameAwsRunShellScript: RunShellScriptFactory{},
	appconfig.PluginNameDomainJoin:        DomainJoinFactory{},
}

type RunShellScriptFactory struct {
}

func (f RunShellScriptFactory) Create(context context.T) (runpluginutil.T, error) {
	switch strings.ToLower(context.AppConfig().RunScriptExecutorName) {
	case "endpoint":
		return endpoint.NewEndpointPlugin(context)
	case "shell":
		return runscript.NewRunShellPlugin(context)
	case "dual":
		return endpoint.NewMuxPlugin(context)
	default:
		return runscript.NewRunShellPlugin(context)
	}
}

type DomainJoinFactory struct {
}

func (f DomainJoinFactory) Create(context context.T) (runpluginutil.T, error) {
	return domainjoin.NewPlugin(context)
}

// loadPlatformDependentPlugins registers platform dependent plugins
func loadPlatformDependentPlugins(context context.T) runpluginutil.PluginRegistry {
	var workerPlugins = runpluginutil.PluginRegistry{}
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
