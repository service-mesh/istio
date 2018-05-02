
//  Copyright 2018 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package showcase

import (
	"istio.io/istio/pkg/test"
	"istio.io/istio/pkg/test/environment"
	"testing"
)

var cfg = ""

// Reimplement TestSvc2Svc in a_simple-1_test.go
func TestSvcLoading(t *testing.T) {
	// This Requires statement should ensure that all elements are in runnig states
	test.Requires(t, denpendency.FortioApps, dependency.Pilot)

	env := test.GetEnvironment(t)
	env.Configure(cfg)

	apps := env.GetFortioApp("app=echosrv")

	path := "/echo"
	arg := "load -qps 0 -t 10s "
	// Test Loading
	for _, app := range apps {
		if _, err := app.CallFotio(arg, path); err != nil {
			t.Fatal("Failed to run fortio %s.", err)
		}
	}
}