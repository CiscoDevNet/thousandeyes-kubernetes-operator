## ThousandEyes Operator
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)

ThousandEyes Operator is a Kubernetes operator used to manage ThousandEyes [Tests](https://developer.thousandeyes.com/v6/tests/) deployed via the Kubernetes cluster.
It is built using the [Operator SDK](https://github.com/operator-framework/operator-sdk), which is part of the [Operator Framework](https://github.com/operator-framework/).

## Documentation
* [Benefits](#benefits)
* [Supported Features](#supported-features)
* [Use ThousandEyes Operator](#use-thousandeyes-operator)
  * [Define a Kubernetes Custom Resource](#define-a-kubernetes-custom-resource)
  * [Annotate a Kubernetes Ingress or Service resource](#annotate-a-kubernetes-ingress-or-service-resource)
* [Quick Start](#quick-start)
  * [Prerequisites](#prerequisites)
  * [Install ThousandEyes Operator](#install-thousandeyes-operator)
  * [Run Tests With Custom Resource](#1-run-tests-with-custom-resource)
    * [Run a HTTP Server Test](#run-a-http-server-test)
    * [Run a Page Load Test](#run-a-page-load-test)
    * [Run a Web Transactions Test](#run-a-web-transactions-test)
  * [Run Tests With Kubernetes Ingress Resource](#2-run-tests-with-kubernetes-ingress-resource)
    * [Install Ingress Controller Locally](#install-ingress-controller-locally)
    * [Run a HTTP Server Test](#run-a-http-server-test-1)
    * [Run a Page Load Test](#run-a-page-load-test-1)
    * [Run a Web Transactions Test](#run-a-web-transactions-test-1)

## Benefits
Here are the benefits of building this operator:
* Automate ThousandEyes test operation
* Define ThousandEyes tests as Code
* Manage ThousandEyes tests as K8s resources in a cloud native way
* Ready for DevOps

## Supported Features
At this stage, ThousandEyes Operatator supports creating, updating and deleting the following three test types.
- HTTP Server Test
- Page Load Test
- Web Transacations Test

We will support more types and features in the future.

## Use ThousandEyes Operator
There are two ways to run a ThousandEyes test with the operator:
- Define a Kubernetes Custom Resource
- Annotate a Kubernetes Ingress or Service resource

### Define a Kubernetes Custom Resource

We can define a custom resource(CR) following its Custom Resource Definition (CRD) respectively:

* [HTTP Server Test CRD](./config/crd/bases/thousandeyes.devnet.cisco.com_httpservertests.yaml) ( [CR Sample](./config/samples/devnet_v1alpha1_httpservertest.yaml) )
* [Page Load Test CRD](./config/crd/bases/thousandeyes.devnet.cisco.com_pageloadtests.yaml) ( [CR Sample](./config/samples/devnet_v1alpha1_pageloadtest.yaml) )
* [Web Transactions Test CRD](./config/crd/bases/thousandeyes.devnet.cisco.com_webtransactiontests.yaml) ( [CR Sample](./config/samples/devnet_v1alpha1_webtransactiontest.yaml) )

These fields in Spec could also be found in [ThousandEyes Test Metadata](https://developer.thousandeyes.com/v6/tests/#/test_metadata).

### Annotate a Kubernetes Ingress or Service resource

We can add the following annotations in Ingress or Service resource:

* **thousandeyes.devnet.cisco.com/test-type**: test type
* **thousandeyes.devnet.cisco.com/test-url**: target url
* **thousandeyes.devnet.cisco.com/test-spec**: test settings
* **thousandeyes.devnet.cisco.com/test-script**: transaction script for web transactions test

Below are the Ingress samples for your reference:

* HTTP Server Test: [Case1](./config/samples/ingress/ingress_httpserver_default_settings.yaml)   [Case2](./config/samples/ingress/ingress_httpserver_specific_settings.yaml)   [Case3](./config/samples/ingress/ingress_httpserver_removal.yaml)
* Page Load Test: [Case1](./config/samples/ingress/ingress_pageload_default_settings.yaml)   [Case2](./config/samples/ingress/ingress_pageload_specific_settings.yaml)   [Case3](./config/samples/ingress/ingress_pageload_removal.yaml)
* Web Transactions Test: [Case1](./config/samples/ingress/ingress_webtransactions_default_settings.yaml)   [Case2](./config/samples/ingress/ingress_webtransactions_specific_settings.yaml)   [Case3](./config/samples/ingress/ingress_webtransactions_removal.yaml)

## Quick Start

### Prerequisites

ThousandEyes Operator requires a Kubernetes cluster of version `>=1.18.0`. If you have just started with Operators, it is highly recommended to use latest version of Kubernetes.

### Install ThousandEyes Operator

1. Clone the project
   ```
   git clone https://github.com/CiscoDevNet/thousandeyes-operator.git
   cd thousandeyes-operator
   ```

2. Get the Oauth Bearer Token from [ThousandEyes dashboard](https://app.thousandeyes.com/login)

   ![Oauth Bearer Token](./docs/thousandeyes_token.gif)

3. Update the OAuth Bearer Token

   Encode the token in base64
   ```
   echo -n "YOUR_OAUTH_BEARER_TOKEN" | base64
   WU9VUl9PQVVUSF9CRUFSRVJfVE9LRU4=
   ```

   Modify OAuthBeaerToken (base64 encoded) in [thousandeyes_operator.yaml](config/deploy/thousandeyes_operator.yaml#L7)

4. Install the operator
   ```
   kubectl apply -f config/deploy/thousandeyes_operator.yaml
   ```

5. Verify installation status

   i. Check the ThousandEyes Operator pod status
   ```
   kubectl get pods | grep thousandeyes
     NAME                                                 READY   STATUS    RESTARTS   AGE
     devnet-thousandeyes-operator-564b5d75d-jllzk         1/1     Running   0          108s
   ```
   ii. Check the ThousandEyes CRD status
   ```
   kubectl get crd | grep thousandeyes
     NAME                                                  CREATED AT
     annotationmonitorings.thousandeyes.devnet.cisco.com   2021-07-07T15:44:40Z
     httpservertests.thousandeyes.devnet.cisco.com         2021-07-07T15:44:41Z
     pageloadtests.thousandeyes.devnet.cisco.com           2021-07-07T15:44:42Z
     webtransactiontests.thousandeyes.devnet.cisco.com     2021-07-07T15:44:44Z 
   ```

## 1. Run Tests with Custom Resource

### Run a HTTP Server Test
1. Create a HTTP Server test
    ```
    kubectl apply -f config/samples/devnet_v1alpha1_httpservertest.yaml
    ```
2. Update the settings of the HTTP Server test
  
   Modify the fields specified by [HTTP Server Test CR](./config/samples/devnet_v1alpha1_httpservertest.yaml) and redeploy.
    ```
    kubectl apply -f config/samples/devnet_v1alpha1_httpservertest.yaml
    ```
3. Delete the HTTP Server test
    ```
    kubectl delete -f config/samples/devnet_v1alpha1_httpservertest.yaml
    ```

### Run a Page Load Test
1. Create a Page Load test
    ```
    kubectl apply -f config/samples/devnet_v1alpha1_pageloadtest.yaml
    ```
2. Update the settings of the Page Load test

   Modify the fields specified by [Page Load Test CR](./config/samples/devnet_v1alpha1_pageloadtest.yaml) and redeploy.
    ```
    kubectl apply -f config/samples/devnet_v1alpha1_pageloadtest.yaml
    ```
3. Delete the Page Load test
    ```
    kubectl delete -f config/samples/devnet_v1alpha1_pageloadtest.yaml
    ```

### Run a Web Transactions Test
1. Create a Web Transaction test
   ```
   kubectl apply -f config/samples/devnet_v1alpha1_webtransactiontest.yaml
   ```
2. Update the settings of the Web Transaction test

   Modify the fields specified by [Web Transaction Test CR](./config/samples/devnet_v1alpha1_webtransactiontest.yaml) and redeploy
    ```
    kubectl apply -f config/samples/devnet_v1alpha1_webtransactiontest.yaml
    ```
3. Delete the Web Transaction test
    ```
    kubectl delete -f config/samples/devnet_v1alpha1_webtransactiontest.yaml
    ```
## 2. Run Tests with Kubernetes Ingress Resource

We will take Ingress as an example, it is also applied to Service.

### Install Ingress Controller Locally

There are multiple Ingress controllers, we will use the Nginx Ingress Controller as an instance.

1. Install the Nginx Ingress Controller
   ```
   kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
   ```
2. Check the Ingress pod status
   ```
   kubectl get pods -A | grep ingress-nginx
     ingress-nginx   ingress-nginx-admission-create-2n62c        0/1     Completed   0          66s
     ingress-nginx   ingress-nginx-admission-patch-fwnlg         0/1     Completed   1          66s
     ingress-nginx   ingress-nginx-controller-68649d49b8-62zvc   1/1     Running     0          66s
   ```

### Run a HTTP Server Test
#### 1. Use the Default Settings:
1. Create a HTTP Server Test
   ```
   kubectl apply -f config/samples/ingress/ingress_httpserver_default_settings.yaml
   ```
2. Update the settings of the HTTP Server test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_httpserver_default_settings.yaml#L7) and redeploy.
    ```
    kubectl apply -f config/samples/ingress/ingress_httpserver_default_settings.yaml
    ```
3. Delete the HTTP Server test
    ```
    kubectl apply -f config/samples/ingress/ingress_httpserver_removal.yaml
    ```
#### 2. Use the Specific Settings:
1. Create a HTTP Server Test
   ```
   kubectl apply -f config/samples/ingress/ingress_httpserver_specific_settings.yaml
   ```
2. Update the settings of the HTTP Server test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_httpserver_specific_settings.yaml#L7) and redeploy.
   ```
   kubectl apply -f config/samples/ingress/ingress_httpserver_specific_settings.yaml
   ```
3. Delete the HTTP Server test
   ```
   kubectl apply -f config/samples/ingress/ingress_httpserver_removal.yaml
   ```
### Run a Page Load Test
#### 1. Use the Default Settings:
1. Create a Page Load Test
   ```
   kubectl apply -f config/samples/ingress/ingress_pageload_default_settings.yaml
   ```
2. Update the settings of the Page Load test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_pageload_default_settings.yaml#L7) and redeploy.
    ```
    kubectl apply -f config/samples/ingress/ingress_pageload_default_settings.yaml
    ```
3. Delete the Page Load test
    ```
    kubectl apply -f config/samples/ingress/ingress_pageload_removal.yaml
    ```
#### 2. Use the Specific Settings:
1. Create a Page Load Test
   ```
   kubectl apply -f config/samples/ingress/ingress_pageload_specific_settings.yaml
   ```
2. Update the settings of the Page Load test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_pageload_specific_settings.yaml#L7) and redeploy.
   ```
   kubectl apply -f config/samples/ingress/ingress_pageload_specific_settings.yaml
   ```
3. Delete the Page Load test
   ```
   kubectl apply -f config/samples/ingress/ingress_pageload_removal.yaml
   ```
### Run a Web Transactions Test
#### 1. Use the Default Settings:
1. Create a Web Transactions Test
   ```
   kubectl apply -f config/samples/ingress/ingress_webtransactions_default_settings.yaml
   ```
2. Update the settings of the Web Transactions test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_webtransactions_default_settings.yaml#L7) and redeploy.
    ```
    kubectl apply -f config/samples/ingress/ingress_webtransactions_default_settings.yaml
    ```
3. Delete the Web Transactions test
    ```
    kubectl apply -f config/samples/ingress/ingress_webtransactions_removal.yaml
    ```
#### 2. Use the Specific Settings:
1. Create a Web Transactions Test
   ```
   kubectl apply -f config/samples/ingress/ingress_webtransactions_specific_settings.yaml
   ```
2. Update the settings of the Web Transactions test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_webtransactions_specific_settings.yaml#L7) and redeploy.
   ```
   kubectl apply -f config/samples/ingress/ingress_webtransactions_specific_settings.yaml
   ```
3. Delete the Web Transactions test
   ```
   kubectl apply -f config/samples/ingress/ingress_webtransactions_removal.yaml
   ```

## References
1. [ThousandEyes Getting Started](https://docs.thousandeyes.com/product-documentation/getting-started)
2. [ThousandEyes Test MetaData](https://developer.thousandeyes.com/v6/tests/#/test_metadata)












