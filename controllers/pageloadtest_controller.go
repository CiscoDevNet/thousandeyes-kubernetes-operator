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
	"github.com/go-logr/logr"
	"github.com/william20111/go-thousandeyes"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	devnetv1alpha1 "wwwin-github.cisco.com/DevNet/thousandeyes-operator/api/v1alpha1"
)

const pageloadFinalizer = "thousandeyes.devnet.cisco.com.pageload.finalizer"

// PageLoadTestReconciler reconciles a PageLoadTest object
type PageLoadTestReconciler struct {
	client.Client
	Log                logr.Logger
	Scheme             *runtime.Scheme
	ThousandEyesClient *thousandeyes.Client
}

//+kubebuilder:rbac:groups=thousandeyes.devnet.cisco.com,resources=pageloadtests,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=thousandeyes.devnet.cisco.com,resources=pageloadtests/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=thousandeyes.devnet.cisco.com,resources=pageloadtests/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PageLoadTest object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *PageLoadTestReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("pageloadtest", req.NamespacedName)

	// fetch the instance
	pl := &devnetv1alpha1.PageLoadTest{}
	err := r.Get(ctx, req.NamespacedName, pl)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	tests, _ := r.ThousandEyesClient.GetTests()
	if tests != nil {
		for _, test := range *tests {
			if test.TestName == pl.Name {
				pl.Spec.PageLoad.TestID = test.TestID
				break
			}
		}
	}
	// check if the thousandeyes-test instance is marked to be deleted, which is
	// indicated by the deletion timestamp being set
	isTeTestMarkedToBeDeleted := pl.GetDeletionTimestamp() != nil
	if isTeTestMarkedToBeDeleted {
		if controllerutil.ContainsFinalizer(pl, pageloadFinalizer) {
			// delete thousandeyes test from server
			if pl.Spec.PageLoad.TestID != 0 {
				err = r.ThousandEyesClient.DeletePageLoad(pl.Spec.PageLoad.TestID)
				if err != nil {
					return ctrl.Result{}, err
				}
			}
			controllerutil.RemoveFinalizer(pl, pageloadFinalizer)
			err := r.Update(ctx, pl)
			if err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !controllerutil.ContainsFinalizer(pl, pageloadFinalizer) {
		controllerutil.AddFinalizer(pl, pageloadFinalizer)
		err := r.Update(ctx, pl)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	data := PageLoad(pl.Spec.PageLoad)

	agents, err := r.ThousandEyesClient.GetAgents()
	if err != nil {
		return ctrl.Result{}, err
	}
	data.Agents = Agents(pl.Spec.Agents, *agents)

	if len(pl.Spec.AlertRules) != 0 {
		alertRules, err := r.ThousandEyesClient.GetAlertRules()
		if err != nil {
			return ctrl.Result{}, err
		}
		data.AlertRules = AlertRules(pl.Spec.AlertRules, *alertRules)
	}

	if pl.Spec.PageLoad.TestID != 0 {
		//check if the test needs to be updated
		pageLoad, err := r.ThousandEyesClient.GetPageLoad(pl.Spec.PageLoad.TestID)
		if err != nil {
			return ctrl.Result{}, err
		}
		if !ComparePageLoad(pl.Spec.PageLoad, *pageLoad) {
			_, err := r.ThousandEyesClient.UpdatePageLoad(pl.Spec.PageLoad.TestID, data)
			if err != nil {
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, nil
	}
	data.TestName = pl.Name
	_, err = r.ThousandEyesClient.CreatePageLoad(data)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PageLoadTestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devnetv1alpha1.PageLoadTest{}).
		WithOptions(controller.Options{
			MaxConcurrentReconciles: maxConcurrentReconciles,
		}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
