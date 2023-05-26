/*
Copyright 2023.

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

package controller

import (
	"context"

	kubeerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	v1alpha1 "github.com/oceanbase/ob-operator/api/v1alpha1"
	"github.com/oceanbase/ob-operator/pkg/resource"
)

// OBServerReconciler reconciles a OBServer object
type OBServerReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=cloud.oceanbase.com,resources=observers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cloud.oceanbase.com,resources=observers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cloud.oceanbase.com,resources=observers/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=core,resources=pods/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=persistentvolumeclaims/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=core,resources=persistentvolumes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=persistentvolumes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups="",resources=events,verbs=create;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the OBServer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *OBServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	observer := &v1alpha1.OBServer{}
	err := r.Client.Get(ctx, req.NamespacedName, observer)
	if err != nil {
		logger.Error(err, "get observer error")
		if kubeerrors.IsNotFound(err) {
			// observer not found, just return
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}
	logger.Info("reconcile observer:", "spec", observer.Spec, "status", observer.Status)

	// TODO add finalizers

	// create observer manager
	observerManager := &resource.OBServerManager{
		Ctx:      ctx,
		OBServer: observer,
		Client:   r.Client,
		Recorder: r.Recorder,
		Logger:   &logger,
	}
	coordinator := resource.NewCoordinator(observerManager, &logger)
	err = coordinator.Coordinate()
	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *OBServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.OBServer{}).
		Complete(r)
}