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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// OperatorConfigurationApplyConfiguration represents an declarative configuration of the OperatorConfiguration type for use
// with apply.
type OperatorConfigurationApplyConfiguration struct {
	CustomImage           *string                                            `json:"customImage,omitempty"`
	CustomImagePullPolicy *string                                            `json:"customImagePullPolicy,omitempty"`
	Namespace             *string                                            `json:"namespace,omitempty"`
	Global                *bool                                              `json:"global,omitempty"`
	ClusterType           *string                                            `json:"clusterType,omitempty"`
	Health                *OperatorHealthConfigurationApplyConfiguration     `json:"health,omitempty"`
	Monitoring            *OperatorMonitoringConfigurationApplyConfiguration `json:"monitoring,omitempty"`
	Debugging             *OperatorDebuggingConfigurationApplyConfiguration  `json:"debugging,omitempty"`
	Tolerations           []string                                           `json:"tolerations,omitempty"`
	NodeSelectors         []string                                           `json:"nodeSelectors,omitempty"`
	ResourcesRequirements []string                                           `json:"resourcesRequirements,omitempty"`
	EnvVars               []string                                           `json:"envVars,omitempty"`
}

// OperatorConfigurationApplyConfiguration constructs an declarative configuration of the OperatorConfiguration type for use with
// apply.
func OperatorConfiguration() *OperatorConfigurationApplyConfiguration {
	return &OperatorConfigurationApplyConfiguration{}
}

// WithCustomImage sets the CustomImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CustomImage field is set to the value of the last call.
func (b *OperatorConfigurationApplyConfiguration) WithCustomImage(value string) *OperatorConfigurationApplyConfiguration {
	b.CustomImage = &value
	return b
}

// WithCustomImagePullPolicy sets the CustomImagePullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CustomImagePullPolicy field is set to the value of the last call.
func (b *OperatorConfigurationApplyConfiguration) WithCustomImagePullPolicy(value string) *OperatorConfigurationApplyConfiguration {
	b.CustomImagePullPolicy = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *OperatorConfigurationApplyConfiguration) WithNamespace(value string) *OperatorConfigurationApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithGlobal sets the Global field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Global field is set to the value of the last call.
func (b *OperatorConfigurationApplyConfiguration) WithGlobal(value bool) *OperatorConfigurationApplyConfiguration {
	b.Global = &value
	return b
}

// WithClusterType sets the ClusterType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterType field is set to the value of the last call.
func (b *OperatorConfigurationApplyConfiguration) WithClusterType(value string) *OperatorConfigurationApplyConfiguration {
	b.ClusterType = &value
	return b
}

// WithHealth sets the Health field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Health field is set to the value of the last call.
func (b *OperatorConfigurationApplyConfiguration) WithHealth(value *OperatorHealthConfigurationApplyConfiguration) *OperatorConfigurationApplyConfiguration {
	b.Health = value
	return b
}

// WithMonitoring sets the Monitoring field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Monitoring field is set to the value of the last call.
func (b *OperatorConfigurationApplyConfiguration) WithMonitoring(value *OperatorMonitoringConfigurationApplyConfiguration) *OperatorConfigurationApplyConfiguration {
	b.Monitoring = value
	return b
}

// WithDebugging sets the Debugging field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Debugging field is set to the value of the last call.
func (b *OperatorConfigurationApplyConfiguration) WithDebugging(value *OperatorDebuggingConfigurationApplyConfiguration) *OperatorConfigurationApplyConfiguration {
	b.Debugging = value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *OperatorConfigurationApplyConfiguration) WithTolerations(values ...string) *OperatorConfigurationApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithNodeSelectors adds the given value to the NodeSelectors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NodeSelectors field.
func (b *OperatorConfigurationApplyConfiguration) WithNodeSelectors(values ...string) *OperatorConfigurationApplyConfiguration {
	for i := range values {
		b.NodeSelectors = append(b.NodeSelectors, values[i])
	}
	return b
}

// WithResourcesRequirements adds the given value to the ResourcesRequirements field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ResourcesRequirements field.
func (b *OperatorConfigurationApplyConfiguration) WithResourcesRequirements(values ...string) *OperatorConfigurationApplyConfiguration {
	for i := range values {
		b.ResourcesRequirements = append(b.ResourcesRequirements, values[i])
	}
	return b
}

// WithEnvVars adds the given value to the EnvVars field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EnvVars field.
func (b *OperatorConfigurationApplyConfiguration) WithEnvVars(values ...string) *OperatorConfigurationApplyConfiguration {
	for i := range values {
		b.EnvVars = append(b.EnvVars, values[i])
	}
	return b
}