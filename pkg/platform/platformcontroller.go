/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package platform

import (
	"fmt"
	"os"
)

const (
	PlatformControllerWatchNamespaceEnvVariable = "WATCH_NAMESPACE"
	platformControllerNamespaceEnvVariable      = "NAMESPACE"
	platformControllerPodNameEnvVariable        = "POD_NAME"
)

const PlatformControllerLockName = "camel-k-platform-controller-lock"

var PlatformControllerImage string

// GetPlatformControllerNamespace returns the namespace where the current platform controller is located (if set).
func GetPlatformControllerNamespace() string {
	if podNamespace, envSet := os.LookupEnv(platformControllerNamespaceEnvVariable); envSet {
		return podNamespace
	}
	return ""
}

// GetPlatformControllerPodName returns the pod that is running the current platform controller (if any).
func GetPlatformControllerPodName() string {
	if podName, envSet := os.LookupEnv(platformControllerPodNameEnvVariable); envSet {
		return podName
	}
	return ""
}

// GetPlatformControllerLockName returns the name of the lock lease that is electing a leader on the particular namespace.
func GetPlatformControllerLockName(platformControllerID string) string {
	return fmt.Sprintf("camel-k-platform-controller-%s-lock", platformControllerID)
}