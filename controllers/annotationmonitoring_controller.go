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
	"encoding/json"
	"github.com/go-logr/logr"
	"github.com/william20111/go-thousandeyes"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"strconv"
	devnetv1alpha1 "wwwin-github.cisco.com/DevNet/thousandeyes-operator/api/v1alpha1"
)

const (
	annotationTestType   = "thousandeyes.devnet.cisco.com/test-type"
	annotationTestURL    = "thousandeyes.devnet.cisco.com/test-url"
	annotationTestScript = "thousandeyes.devnet.cisco.com/test-script"
	annotationTestSpec   = "thousandeyes.devnet.cisco.com/test-spec"
)

// AnnotationMonitoringReconciler reconciles a AnnotationMonitoring object
type AnnotationMonitoringReconciler struct {
	client.Client
	Log                logr.Logger
	Scheme             *runtime.Scheme
	ThousandEyesClient *thousandeyes.Client
}

//+kubebuilder:rbac:groups=networking,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking,resources=ingresses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=networking,resources=ingresses/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AnnotationMonitoring object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *AnnotationMonitoringReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("annotationmonitoring", req.NamespacedName)

	var testName string
	annotations := map[string]string{}
	ingress := &v1.Ingress{}
	err := r.Get(ctx, req.NamespacedName, ingress)
	if err != nil {
		if errors.IsNotFound(err) {
			service := &corev1.Service{}
			err = r.Get(ctx, req.NamespacedName, service)
			if err != nil {
				if errors.IsNotFound(err) {
					return ctrl.Result{}, nil
				}
				return ctrl.Result{}, err
			}
			testName = service.Name
			annotations = service.Annotations
		} else {
			return ctrl.Result{}, err
		}
	} else {
		testName = ingress.Name
		annotations = ingress.Annotations
	}

	if testType, ok := annotations[annotationTestType]; ok {
		tests, err := r.ThousandEyesClient.GetTests()
		if err != nil {
			return ctrl.Result{}, err
		}
		agents, err := r.ThousandEyesClient.GetAgents()
		if err != nil {
			return ctrl.Result{}, err
		}

		switch testType {
		case none:
			for _, test := range *tests {
				if test.TestName == testName {
					switch test.Type {
					case httpserver:
						err = r.ThousandEyesClient.DeleteHTTPServer(test.TestID)
						if err != nil {
							return ctrl.Result{}, err
						}
					case pageload:
						err = r.ThousandEyesClient.DeletePageLoad(test.TestID)
						if err != nil {
							return ctrl.Result{}, err
						}
					case webtransactions:
						err = r.ThousandEyesClient.DeleteWebTransaction(test.TestID)
						if err != nil {
							return ctrl.Result{}, err
						}
					}
					return ctrl.Result{}, nil
				}
			}
		case httpserver:
			httpSpec := &devnetv1alpha1.HTTPServerTestSpec{}
			if specStr, ok := annotations[annotationTestSpec]; ok {
				err = json.Unmarshal([]byte(specStr), httpSpec)
				if err != nil {
					return ctrl.Result{}, err
				}
			} else if url, ok := annotations[annotationTestURL]; ok {
				httpSpec.URL = url
				interval, err := strconv.Atoi(os.Getenv("DEFAULT_INTERVAL"))
				if err != nil {
					return ctrl.Result{}, err
				}
				httpSpec.Interval = interval
				httpSpec.Agents = DefaultAgents()
			} else {
				return ctrl.Result{}, nil
			}
			data := HTTPServer(httpSpec.HTTPServer)
			data.Agents = Agents(httpSpec.Agents, *agents)
			if len(httpSpec.AlertRules) != 0 {
				alertRules, err := r.ThousandEyesClient.GetAlertRules()
				if err != nil {
					return ctrl.Result{}, err
				}
				data.AlertRules = AlertRules(httpSpec.AlertRules, *alertRules)
			}
			for _, test := range *tests {
				if test.TestName == testName {
					httpSpec.TestID = test.TestID
					break
				}
			}
			if httpSpec.HTTPServer.TestID != 0 {
				//check if the test needs to be updated
				httpServer, err := r.ThousandEyesClient.GetHTTPServer(httpSpec.HTTPServer.TestID)
				if err != nil {
					return ctrl.Result{}, err
				}
				if !CompareHTTPServer(httpSpec.HTTPServer, *httpServer) {
					_, err := r.ThousandEyesClient.UpdateHTTPServer(httpSpec.HTTPServer.TestID, data)
					if err != nil {
						return ctrl.Result{}, err
					}
					return ctrl.Result{}, nil
				}
				return ctrl.Result{}, nil
			}
			data.TestName = testName
			_, err = r.ThousandEyesClient.CreateHTTPServer(data)
			if err != nil {
				return ctrl.Result{}, err
			}
		case pageload:
			pageloadSpec := &devnetv1alpha1.PageLoadTestSpec{}
			if specStr, ok := annotations[annotationTestSpec]; ok {
				err = json.Unmarshal([]byte(specStr), pageloadSpec)
				if err != nil {
					return ctrl.Result{}, err
				}
			} else if url, ok := annotations[annotationTestURL]; ok {
				pageloadSpec.URL = url
				interval, err := strconv.Atoi(os.Getenv("DEFAULT_INTERVAL"))
				if err != nil {
					return ctrl.Result{}, err
				}
				httpInterval, err := strconv.Atoi(os.Getenv("DEFAULT_HTTP_INTERVAL"))
				if err != nil {
					return ctrl.Result{}, err
				}
				pageloadSpec.Interval = interval
				pageloadSpec.HTTPInterval = httpInterval
				pageloadSpec.Agents = DefaultAgents()
			} else {
				return ctrl.Result{}, nil
			}
			data := PageLoad(pageloadSpec.PageLoad)
			data.Agents = Agents(pageloadSpec.Agents, *agents)
			if len(pageloadSpec.AlertRules) != 0 {
				alertRules, err := r.ThousandEyesClient.GetAlertRules()
				if err != nil {
					return ctrl.Result{}, err
				}
				data.AlertRules = AlertRules(pageloadSpec.AlertRules, *alertRules)
			}
			for _, test := range *tests {
				if test.TestName == testName {
					pageloadSpec.TestID = test.TestID
					break
				}
			}
			if pageloadSpec.PageLoad.TestID != 0 {
				//check if the test needs to be updated
				pageload, err := r.ThousandEyesClient.GetPageLoad(pageloadSpec.PageLoad.TestID)
				if err != nil {
					return ctrl.Result{}, err
				}
				if !ComparePageLoad(pageloadSpec.PageLoad, *pageload) {
					_, err := r.ThousandEyesClient.UpdatePageLoad(pageloadSpec.PageLoad.TestID, data)
					if err != nil {
						return ctrl.Result{}, err
					}
					return ctrl.Result{}, nil
				}
				return ctrl.Result{}, nil
			}
			data.TestName = testName
			_, err = r.ThousandEyesClient.CreatePageLoad(data)
			if err != nil {
				return ctrl.Result{}, err
			}
		case webtransactions:
			webtransactionSpec := &devnetv1alpha1.WebTransactionTestSpec{}
			if specStr, ok := annotations[annotationTestSpec]; ok {
				err = json.Unmarshal([]byte(specStr), webtransactionSpec)
				if err != nil {
					return ctrl.Result{}, err
				}
			} else if url, ok := annotations[annotationTestURL]; ok {
				webtransactionSpec.URL = url
				interval, err := strconv.Atoi(os.Getenv("DEFAULT_INTERVAL"))
				if err != nil {
					return ctrl.Result{}, err
				}
				webtransactionSpec.Interval = interval
				webtransactionSpec.Agents = DefaultAgents()
			} else {
				return ctrl.Result{}, nil
			}
			if script, ok := annotations[annotationTestScript]; ok {
				webtransactionSpec.TransactionScript = script
			} else {
				return ctrl.Result{}, nil
			}
			data := WebTransaction(webtransactionSpec.WebTransaction)
			data.Agents = Agents(webtransactionSpec.Agents, *agents)
			if len(webtransactionSpec.AlertRules) != 0 {
				alertRules, err := r.ThousandEyesClient.GetAlertRules()
				if err != nil {
					return ctrl.Result{}, err
				}
				data.AlertRules = AlertRules(webtransactionSpec.AlertRules, *alertRules)
			}
			for _, test := range *tests {
				if test.TestName == testName {
					webtransactionSpec.TestID = test.TestID
					break
				}
			}
			if webtransactionSpec.WebTransaction.TestID != 0 {
				//check if the test needs to be updated
				webtransactions, err := r.ThousandEyesClient.GetWebTransaction(webtransactionSpec.WebTransaction.TestID)
				if err != nil {
					return ctrl.Result{}, err
				}
				if !CompareWebTransaction(webtransactionSpec.WebTransaction, *webtransactions) {
					_, err := r.ThousandEyesClient.UpdateWebTransaction(webtransactionSpec.WebTransaction.TestID, data)
					if err != nil {
						return ctrl.Result{}, err
					}
					return ctrl.Result{}, nil
				}
				return ctrl.Result{}, nil
			}
			data.TestName = testName
			_, err = r.ThousandEyesClient.CreateWebTransaction(data)
			if err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AnnotationMonitoringReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devnetv1alpha1.AnnotationMonitoring{}).
		Watches(&source.Kind{Type: &v1.Ingress{}}, &handler.EnqueueRequestForObject{}).
		Watches(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForObject{}).
		Complete(r)
}
