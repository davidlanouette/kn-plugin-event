/*
Copyright 2022 The Knative Authors

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

package manifest

import (
	"context"

	kubeclient "knative.dev/pkg/client/injection/kube/client"
	pkgsecurity "knative.dev/pkg/test/security"
	"knative.dev/reconciler-test/pkg/environment"
	"knative.dev/reconciler-test/pkg/feature"
	"knative.dev/reconciler-test/pkg/k8s"
)

// PodSecurityCfgFn returns a function for configuring security context for Pod, depending
// on security settings of the enclosing namespace.
func PodSecurityCfgFn(ctx context.Context, t feature.T) CfgFn {
	namespace := environment.FromContext(ctx).Namespace()
	restrictedMode, err := pkgsecurity.IsRestrictedPodSecurityEnforced(ctx, kubeclient.Get(ctx), namespace)
	if err != nil {
		t.Fatalf("Error while checking restricted pod security mode for namespace %s", namespace)
	}
	if restrictedMode {
		return k8s.WithDefaultPodSecurityContext
	}
	return func(map[string]interface{}) {}
}

// WithAnnotations returns a function for configuring annototations of the resource
func WithAnnotations(annotations map[string]interface{}) CfgFn {
	return func(cfg map[string]interface{}) {
		if annotations != nil {
			cfg["annotations"] = annotations
		}
	}
}

// WithLabels returns a function for configuring labels of the resource
func WithLabels(labels map[string]string) CfgFn {
	return func(cfg map[string]interface{}) {
		if labels != nil {
			cfg["labels"] = labels
		}
	}
}
