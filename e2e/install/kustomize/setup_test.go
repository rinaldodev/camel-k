//go:build integration
// +build integration

// To enable compilation of this file in Goland, go to "Settings -> Go -> Vendoring & Build Tags -> Custom Tags" and add "integration"

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

package kustomize

import (
	"fmt"
	"os"
	"testing"

	. "github.com/apache/camel-k/v2/e2e/support"
	testutil "github.com/apache/camel-k/v2/e2e/support/util"
	. "github.com/onsi/gomega"
)

const (
	// v1.Build,          v1.Integration
	// v1.IntegrationKit, v1.IntegrationPlatform
	// v1.Kamelet,  v1.Pipe,
	// v1alpha1.Kamelet, v1alpha1.KameletBinding
	ExpectedCRDs = 8

	// camel-k-operator,             camel-k-operator-events,
	// camel-k-operator-knative,     camel-k-operator-leases,
	// camel-k-operator-podmonitors, camel-k-operator-strimzi,
	// camel-k-operator-keda
	ExpectedKubePromoteRoles = 7

	// camel-k-edit
	// camel-k-operator-custom-resource-definitions
	// camel-k-operator-bind-addressable-resolver
	// camel-k-operator-local-registry
	ExpectedKubeClusterRoles = 4

	// camel-k-operator-openshift
	ExpectedOSPromoteRoles = 1

	// camel-k-operator-console-openshift
	ExpectedOSClusterRoles = 1
)

func TestSetupKustomizeBasic(t *testing.T) {
	RegisterTestingT(t)
	makeDir := testutil.MakeTempCopyDir(t, "../../../install")
	os.Setenv("CAMEL_K_TEST_MAKE_DIR", makeDir)

	// Ensure no CRDs are already installed
	UninstallAll()
	Eventually(CRDs()).Should(HaveLen(0))

	// Return the cluster to previous state
	defer Cleanup()

	WithNewTestNamespace(t, func(ns string) {
		namespaceArg := fmt.Sprintf("NAMESPACE=%s", ns)
		ExpectExecSucceed(t, Make("setup-cluster", namespaceArg))
		Eventually(CRDs()).Should(HaveLen(ExpectedCRDs))

		ExpectExecSucceed(t, Make("setup", namespaceArg))

		kpRoles := ExpectedKubePromoteRoles
		opRoles := kpRoles + ExpectedOSPromoteRoles
		Eventually(Role(ns)).Should(Or(HaveLen(kpRoles), HaveLen(opRoles)))

		kcRoles := ExpectedKubeClusterRoles
		ocRoles := kcRoles + ExpectedOSClusterRoles
		Eventually(ClusterRole()).Should(Or(HaveLen(kcRoles), HaveLen(ocRoles)))

		// Tidy up to ensure next test works
		Expect(Kamel("uninstall", "-n", ns).Execute()).To(Succeed())
	})

}

func TestSetupKustomizeGlobal(t *testing.T) {
	makeDir := testutil.MakeTempCopyDir(t, "../../../install")
	os.Setenv("CAMEL_K_TEST_MAKE_DIR", makeDir)

	// Ensure no CRDs are already installed
	UninstallAll()
	Eventually(CRDs()).Should(HaveLen(0))

	// Return the cluster to previous state
	defer Cleanup()

	WithNewTestNamespace(t, func(ns string) {
		namespaceArg := fmt.Sprintf("NAMESPACE=%s", ns)
		ExpectExecSucceed(t, Make("setup-cluster", namespaceArg))
		Eventually(CRDs()).Should(HaveLen(ExpectedCRDs))

		ExpectExecSucceed(t, Make("setup", "GLOBAL=true", namespaceArg))

		Eventually(Role(ns)).Should(HaveLen(0))

		kcpRoles := ExpectedKubeClusterRoles + ExpectedKubePromoteRoles
		ocpRoles := kcpRoles + ExpectedOSClusterRoles + ExpectedOSPromoteRoles
		Eventually(ClusterRole()).Should(Or(HaveLen(kcpRoles), HaveLen(ocpRoles)))
	})
}
