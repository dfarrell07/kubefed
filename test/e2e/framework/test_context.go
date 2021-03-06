/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	"flag"
	"fmt"
	"os"
	"time"

	"k8s.io/klog"

	"sigs.k8s.io/kubefed/pkg/controller/util"
)

type TestContextType struct {
	InMemoryControllers             bool
	KubeConfig                      string
	KubeContext                     string
	KubefedSystemNamespace          string
	SingleCallTimeout               time.Duration
	LimitedScope                    bool
	LimitedScopeInMemoryControllers bool
	WaitForFinalization             bool
}

func (t *TestContextType) RunControllers() bool {
	return t.InMemoryControllers
}

var TestContext *TestContextType = &TestContextType{}

func registerFlags(t *TestContextType) {
	flag.BoolVar(&t.InMemoryControllers, "in-memory-controllers", false,
		"Whether kubefed controllers should be started in memory.")
	flag.StringVar(&t.KubeConfig, "kubeconfig", os.Getenv("KUBECONFIG"),
		"Path to kubeconfig containing embedded authinfo.")
	flag.StringVar(&t.KubeContext, "context", "",
		"kubeconfig context to use/override. If unset, will use value from 'current-context'.")
	flag.StringVar(&t.KubefedSystemNamespace, "kubefed-namespace", util.DefaultKubefedSystemNamespace,
		fmt.Sprintf("The namespace the kubefed control plane is deployed in.  If unset, will default to %q.", util.DefaultKubefedSystemNamespace))
	flag.DurationVar(&t.SingleCallTimeout, "single-call-timeout", DefaultSingleCallTimeout,
		fmt.Sprintf("The maximum duration of a single call.  If unset, will default to %v", DefaultSingleCallTimeout))
	flag.BoolVar(&t.LimitedScope, "limited-scope", false, "Whether the kubefed namespace (configurable via --kubefed-namespace) will be the only target for federation.")
	flag.BoolVar(&t.LimitedScopeInMemoryControllers, "limited-scope-in-memory-controllers", true,
		"Whether kubefed controllers started in memory should target only the test namespace.  If debugging cluster-scoped federation outside of a test namespace, this should be set to false.")
	flag.BoolVar(&t.WaitForFinalization, "wait-for-finalization", true,
		"Whether the test suite should wait for finalization before stopping fixtures or exiting.  Setting this to false will speed up test execution but likely result in wedged namespaces and is only recommended for disposeable clusters.")
}

func validateFlags(t *TestContextType) {
	if len(t.KubeConfig) == 0 {
		klog.Fatalf("kubeconfig is required")
	}
	if t.InMemoryControllers {
		klog.Info("in-memory-controllers=true - this will launch the kubefed controllers outside the cluster hosting the kubefed control plane.")
	}
}

func ParseFlags() {
	registerFlags(TestContext)
	flag.Parse()
	validateFlags(TestContext)
}
