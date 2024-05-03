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
//go:build freebsd || linux || netbsd || openbsd
// +build freebsd linux netbsd openbsd

// Package runpluginutil run plugin utility functions without referencing the actually plugin impl packages
package runpluginutil

import (
	"fmt"

	"github.com/aws/amazon-ssm-agent/agent/context"
	"github.com/aws/amazon-ssm-agent/agent/log"
	"github.com/aws/amazon-ssm-agent/agent/platform"
)

// IsPluginSupportedForCurrentPlatform always returns true for plugins that exist for linux because currently there
// are no plugins that are supported on only one distribution or version of linux.
func IsPluginSupportedForCurrentPlatform(context context.T, log log.T, pluginName string) (isKnown bool, isSupported bool, message string) {
	platformName, _ := platform.PlatformName(log)
	platformVersion, _ := platform.PlatformVersion(log)

	if len(context.AppConfig().EnablePlugins) != 0 {
		foundPlugin := false
		for _, enabledPlugin := range context.AppConfig().EnablePlugins {
			if enabledPlugin != pluginName {
				continue
			}
			foundPlugin = true
			break
		}
		if !foundPlugin {
			return false, false, fmt.Sprintf("%s v%s", platformName, platformVersion)
		}
	}

	if _, known := allSessionPlugins[pluginName]; known {
		return known, true, fmt.Sprintf("%s v%s", platformName, platformVersion)
	}
	_, known := allPlugins[pluginName]
	return known, true, fmt.Sprintf("%s v%s", platformName, platformVersion)
}
