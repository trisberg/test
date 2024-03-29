/*

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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testv1alpha1 "github.com/trisberg/test/api/v1alpha1"
)

// ExampleReconciler reconciles a Example object
type ExampleReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=test.example.com,resources=examples,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=test.example.com,resources=examples/status,verbs=get;update;patch

func (r *ExampleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("example", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *ExampleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testv1alpha1.Example{}).
		Complete(r)
}
