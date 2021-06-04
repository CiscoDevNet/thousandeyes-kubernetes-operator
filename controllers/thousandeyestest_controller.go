/*
Copyright 2021.

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
	"github.com/william20111/go-thousandeyes"
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	devnetv1alpha1 "wwwin-github.cisco.com/DevNet/thousandeyes-operator/api/v1alpha1"
)

const teTestFinalizer = "devnet.cisco.com.te.test.finalizer"

// ThousandEyesTestReconciler reconciles a ThousandEyesTest object
type ThousandEyesTestReconciler struct {
	client.Client
	Log                logr.Logger
	Scheme             *runtime.Scheme
	ThousandEyesClient *thousandeyes.Client
}

//+kubebuilder:rbac:groups=devnet.cisco.com,resources=thousandeyestests,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=devnet.cisco.com,resources=thousandeyestests/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=devnet.cisco.com,resources=thousandeyestests/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ThousandEyesTest object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *ThousandEyesTestReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("thousandeyestest", req.NamespacedName)

	// fetch the instance
	te := &devnetv1alpha1.ThousandEyesTest{}
	err := r.Get(ctx, req.NamespacedName, te)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	tests, _ := r.ThousandEyesClient.GetTests()
	if tests != nil {
		for _, test := range *tests {
			if test.TestName == te.Name {
				te.Spec.Metadata.TestID = test.TestID
				break
			}
		}
	}
	// check if the thousandeyes-test instance is marked to be deleted, which is
	// indicated by the deletion timestamp being set
	isTeTestMarkedToBeDeleted := te.GetDeletionTimestamp() != nil
	if isTeTestMarkedToBeDeleted {
		if controllerutil.ContainsFinalizer(te, teTestFinalizer) {
			// delete thousandeyes test from server
			if te.Spec.Metadata.TestID != 0 {
				err = DeleteTest(r.ThousandEyesClient, te.Spec)
				if err != nil {
					return ctrl.Result{}, err
				}
				controllerutil.RemoveFinalizer(te, teTestFinalizer)
				err := r.Update(ctx, te)
				if err != nil {
					return ctrl.Result{}, err
				}
			}
		}
		return ctrl.Result{}, nil
	}

	if !controllerutil.ContainsFinalizer(te, teTestFinalizer) {
		controllerutil.AddFinalizer(te, teTestFinalizer)
		err := r.Update(ctx, te)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	if te.Spec.Metadata.TestID != 0 {
		//check if the test needs to be updated
		pageLoad, err := r.ThousandEyesClient.GetPageLoad(te.Spec.Metadata.TestID)
		if err != nil {
			return ctrl.Result{}, err
		}
		if !EqualMetadata(te.Spec.Metadata, *pageLoad) {
			_, err := UpdateTest(r.ThousandEyesClient, te.Spec)
			if err != nil {
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, nil
		}
	}
	te.Spec.Metadata.TestName = te.Name
	_, err = CreateTest(r.ThousandEyesClient, te.Spec)
	if err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ThousandEyesTestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devnetv1alpha1.ThousandEyesTest{}).
		Complete(r)
}
