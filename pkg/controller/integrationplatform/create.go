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

package integrationplatform

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/apache/camel-k/v2/pkg/apis/camel/v1"
	"github.com/apache/camel-k/v2/pkg/install"
	"github.com/apache/camel-k/v2/pkg/resources"
	"github.com/apache/camel-k/v2/pkg/util/defaults"
)

// NewCreateAction returns the action that creates resources needed by the platform.
func NewCreateAction() Action {
	return &createAction{}
}

type createAction struct {
	baseAction
}

func (action *createAction) Name() string {
	return "create"
}

func (action *createAction) CanHandle(platform *v1.IntegrationPlatform) bool {
	return platform.Status.Phase == v1.IntegrationPlatformPhaseCreating
}

func (action *createAction) Handle(ctx context.Context, platform *v1.IntegrationPlatform) (*v1.IntegrationPlatform, error) {
	paths, err := resources.WithPrefix("/resources/camel-catalog-")
	if err != nil {
		return nil, err
	}

	for _, k := range paths {
		action.L.Infof("Installing camel catalog: %s", k)
		err := install.Resources(ctx, action.client, platform.Namespace, true,
			func(object ctrl.Object) ctrl.Object {
				action.L.Infof("Copying platform annotations to catalog: %s", object.GetName())
				object.SetAnnotations(platform.Annotations)
				return object
			},
			k)
		if err != nil {
			return nil, err
		}
	}

	if defaults.InstallDefaultKamelets() {
		// Kamelet Catalog installed on platform reconciliation for cases where users install a global operator
		if err := install.KameletCatalog(ctx, action.client, platform.Namespace); err != nil {
			return nil, err
		}
	}

	action.L.Infof("Installing Camel K operator: %s", platform.Name)
	err = install.InstallOperator(ctx, action.client, platform.Namespace, platform.Spec.OperatorConfiguration)
	if err != nil {
		action.L.Errorf(err, "Failed to install Camel K operator: %s", platform.Name)
		platform.Status.SetCondition(
			v1.IntegrationPlatformConditionOperatorInstalled,
			corev1.ConditionFalse,
			v1.IntegrationPlatformConditionOperatorInstalledReason,
			fmt.Sprintf("unable to install camel k operator: %s", err.Error()))

		return platform, err
	}

	platform.Status.SetCondition(
		v1.IntegrationPlatformConditionOperatorInstalled,
		corev1.ConditionTrue,
		v1.IntegrationPlatformConditionOperatorInstalledReason,
		"camel k operator installed")

	platform.Status.SetCondition(
		v1.IntegrationPlatformConditionTypeCreated,
		corev1.ConditionTrue,
		v1.IntegrationPlatformConditionCreatedReason,
		"integration platform created")

	platform.Status.Phase = v1.IntegrationPlatformPhaseReady
	return platform, nil
}
