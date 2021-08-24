## ThousandEyes Operator
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)

ThousandEyes Operator is Kubernetes operator used to manage ThousandEyes [Tests](https://developer.thousandeyes.com/v6/tests/) deployed via the Kubernetes cluster.
It is built using the [Operator SDK](https://github.com/operator-framework/operator-sdk), which is part of the [Operator Framework](https://github.com/operator-framework/).

The purpose of creating this operator was to provide an easy operations of ThousandEyes on Kubernetes.

### Supported Features
This operatator supports creating, updating and deleting the following test types.
- HTTP Server Test
- Page Load Test 
- Web Transacation Test

### Prerequisites

ThousandEyes Operator requires a Kubernetes cluster of version `>=1.18.0`. If you have just started with Operators, it is highly recommended to use latest version of Kubernetes.

## Quick Start

### Install ThousandEyes Operator

1. Clone the project
   ```
   $ git clone https://github.com/CiscoDevNet/thousandeyes-operator.git
   $ cd thousandeyes-operator
   ```

2. Get the Oauth Bearer Token from [ThousandEyes dashboard](https://app.thousandeyes.com/login)

   ![Oauth Bearer Token](./docs/thousandeyes_token.gif)

3. Update the OAuth Bearer Token

   Encode the token in base64
   ~~~
   $ echo -n "YOUR_OAUTH_BEARER_TOKEN" | base64
   WU9VUl9PQVVUSF9CRUFSRVJfVE9LRU4=
   ~~~

   Modify OAuthBeaerToken (base64 encoded) in [thousandeyes_operator.yaml](config/deploy/thousandeyes_operator.yaml#L7)

4. Install the operator
   ```
   $ kubectl apply -f config/deploy/thousandeyes_operator.yaml
   ```

5. Verify installation status

   i. Check the ThousandEyes Operator pod status
   ```
   $ kubectl get pods | grep thousandeyes
     NAME                                                 READY   STATUS    RESTARTS   AGE
     devnet-thousandeyes-operator-564b5d75d-jllzk         1/1     Running   0          108s
   ```
   ii. Check the ThousandEyes CRD status
   ```
   $ kubectl get crd | grep thousandeyes
     NAME                                                  CREATED AT
     annotationmonitorings.thousandeyes.devnet.cisco.com   2021-07-07T15:44:40Z
     httpservertests.thousandeyes.devnet.cisco.com         2021-07-07T15:44:41Z
     pageloadtests.thousandeyes.devnet.cisco.com           2021-07-07T15:44:42Z
     webtransactiontests.thousandeyes.devnet.cisco.com     2021-07-07T15:44:44Z 
   ```
There are two ways to run a ThousandEyes test:
- Create a Custom Resource(CR) defined in ThousandEyes Operator
- Create a Kubernetes internal resource with its annotations

We will create the sample tests with the approaches above to make it more clear:

## 1. Run the tests with the ThousandEyes Operator CR

### Run a HTTP Server Test
1. Create a HTTP Server test
    ```
    $ kubectl apply -f config/samples/devnet_v1alpha1_httpservertest.yaml
    ```
2. Update the settings of the HTTP Server test
  
   Modify the fields specified by [HTTP Server Test CR](./config/samples/devnet_v1alpha1_httpservertest.yaml) and redeploy.
    ```
    $ kubectl apply -f config/samples/devnet_v1alpha1_httpservertest.yaml
    ```
3. Delete the HTTP Server test
    ```
    $ kubectl delete -f config/samples/devnet_v1alpha1_httpservertest.yaml
    ```

### Run a Page Load Test
1. Create a Page Load test
    ```
    $ kubectl apply -f config/samples/devnet_v1alpha1_pageloadtest.yaml
    ```
2. Update the settings of the Page Load test

   Modify the fields specified by [Page Load Test CR](./config/samples/devnet_v1alpha1_pageloadtest.yaml) and redeploy.
    ```
    $ kubectl apply -f config/samples/devnet_v1alpha1_pageloadtest.yaml
    ```
3. Delete the Page Load test
    ```
    $ kubectl delete -f config/samples/devnet_v1alpha1_pageloadtest.yaml
    ```

### Run a Web Transaction Test
1. Create a Web Transaction test
   ```
   $ kubectl apply -f config/samples/devnet_v1alpha1_webtransactiontest.yaml
   ```
2. Update the settings of the Web Transaction test

   Modify the fields specified by [Web Transaction Test CR](./config/samples/devnet_v1alpha1_webtransactiontest.yaml) and redeploy
    ```
    $ kubectl apply -f config/samples/devnet_v1alpha1_webtransactiontest.yaml
    ```
3. Delete the Web Transaction test
    ```
    $ kubectl delete -f config/samples/devnet_v1alpha1_webtransactiontest.yaml
    ```
## 2. Run the tests with the kubernetes internal resource

In this scenario, all the settings of the tests could be specified in the following **annotations** in kubernetes internal resource.

Once the resource with the annotations is created, the test will be created as well.

- thousandeyes.devnet.cisco.com/test-type: specify the test type. (**required**)
- thousandeyes.devnet.cisco.com/test-url: specify the target url with its default settings. (**required if not specify thousandeyes.devnet.cisco.com/test-spec**)
- thousandeyes.devnet.cisco.com/test-script: specfy the test script for web transaction test. (**required for web transaction test**)
- thousandeyes.devnet.cisco.com/test-spec: specify the settings of this test. (**required if not specify thousandeyes.devnet.cisco.com/test-url**)

At this point, we support two resources: **Ingress** and **Service**.

Here we take Ingress as an example, it is also applied to Service.

### Install Ingress Controller Locally

There are multiple Ingress controllers, we will use the Nginx Ingress Controller as an instance.

1. Install the Nginx Ingress Controller
   ```
   $ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
   ```
2. Check the Ingress pod status
   ```
   $ kubectl get pods -A | grep ingress-nginx
     ingress-nginx   ingress-nginx-admission-create-2n62c        0/1     Completed   0          66s
     ingress-nginx   ingress-nginx-admission-patch-fwnlg         0/1     Completed   1          66s
     ingress-nginx   ingress-nginx-controller-68649d49b8-62zvc   1/1     Running     0          66s
   ```

### Run a HTTP Server Test (Default Settings)
1. Create a HTTP Server Test
   ```
   $ kubectl apply -f config/samples/ingress/ingress_httpserver_default_settings.yaml
   ```
2. Update the settings of the HTTP Server test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_httpserver_default_settings.yaml#L7) and redeploy.
    ```
    $ kubectl apply -f config/samples/ingress/ingress_httpserver_default_settings.yaml
    ```
3. Delete the HTTP Server test
    ```
    $ kubectl apply -f config/samples/ingress/ingress_httpserver_removal.yaml
    ```
### Run a HTTP Server Test (Specific Settings)  
1. Create a HTTP Server Test
   ```
   $ kubectl apply -f config/samples/ingress/ingress_httpserver_specific_settings.yaml
   ```
2. Update the settings of the HTTP Server test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_httpserver_specific_settings.yaml#L7) and redeploy.
   ```
   $ kubectl apply -f config/samples/ingress/ingress_httpserver_specific_settings.yaml
   ```
3. Delete the HTTP Server test
   ```
   $ kubectl apply -f config/samples/ingress/ingress_httpserver_removal.yaml
   ```
### Run a Page Load Test (Default Settings)
1. Create a Page Load Test
   ```
   $ kubectl apply -f config/samples/ingress/ingress_pageload_default_settings.yaml
   ```
2. Update the settings of the Page Load test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_pageload_default_settings.yaml#L7) and redeploy.
    ```
    $ kubectl apply -f config/samples/ingress/ingress_pageload_default_settings.yaml
    ```
3. Delete the Page Load test
    ```
    $ kubectl apply -f config/samples/ingress/ingress_pageload_removal.yaml
    ```
### Run a Page Load Test (Specific Settings)
1. Create a Page Load Test
   ```
   $ kubectl apply -f config/samples/ingress/ingress_pageload_specific_settings.yaml
   ```
2. Update the settings of the Page Load test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_pageload_specific_settings.yaml#L7) and redeploy.
   ```
   $ kubectl apply -f config/samples/ingress/ingress_pageload_specific_settings.yaml
   ```
3. Delete the Page Load test
   ```
   $ kubectl apply -f config/samples/ingress/ingress_pageload_removal.yaml
   ```
### Run a Web Transactions Test (Default Settings)
1. Create a Web Transactions Test
   ```
   $ kubectl apply -f config/samples/ingress/ingress_webtransactions_default_settings.yaml
   ```
2. Update the settings of the Web Transactions test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_webtransactions_default_settings.yaml#L7) and redeploy.
    ```
    $ kubectl apply -f config/samples/ingress/ingress_webtransactions_default_settings.yaml
    ```
3. Delete the Web Transactions test
    ```
    $ kubectl apply -f config/samples/ingress/ingress_webtransactions_removal.yaml
    ```
### Run a Web Transactions Test (Specific Settings)
1. Create a Web Transactions Test
   ```
   $ kubectl apply -f config/samples/ingress/ingress_webtransactions_specific_settings.yaml
   ```
2. Update the settings of the Web Transactions test

   Modify the annotation specified by [Ingress resource](./config/samples/ingress/ingress_webtransactions_specific_settings.yaml#L7) and redeploy.
   ```
   $ kubectl apply -f config/samples/ingress/ingress_webtransactions_specific_settings.yaml
   ```
3. Delete the Web Transactions test
   ```
   $ kubectl apply -f config/samples/ingress/ingress_webtransactions_removal.yaml
   ```

## References
1. [ThousandEyes Getting Started](https://docs.thousandeyes.com/product-documentation/getting-started)
2. [ThousandEyes Test MetaData](https://developer.thousandeyes.com/v6/tests/#/test_metadata)












