## ThousandEyes Operator
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)

ThousandEyes Operator is Kubernetes operator used to manage ThousandEyes [Tests](https://developer.thousandeyes.com/v6/tests/) deployed via the Kubernetes cluster.
It is built using the [Operator SDK](https://github.com/operator-framework/operator-sdk), which is part of the [Operator Framework](https://github.com/operator-framework/).

The purpose of creating this operator was to provide an easy operations of ThousandEyes on Kubernetes.

### Supported Features
This operatator supports creating, updating and deleting the following test types:
- HTTP Server Test
- Page Load Test 
- Web Transacation Test

### Prerequisites

ThousandEyes Operator requires a Kubernetes cluster of version `>=1.16.0`. If you have just started with Operators, its highly recommended to use latest version of Kubernetes.

## Quick Start

### Install ThousandEyes Operator

1. Clone the project:
```
$ git clone https://github.com/CiscoDevNet/thousandeyes-operator.git
$ cd thousandeyes-operator
```

2. Get the Oauth Bearer Token from [ThousandEyes dashboard](https://app.thousandeyes.com/login):
* go to **Account Settings > Users and Roles > User API Tokens > OAuth Bearer Token** in [ThousandEyes dashboard](https://app.thousandeyes.com/login)
* set the environment variable **THOUSANDEYES_CLIENT_TOKEN** with the **OAuth Bearer Token** in [thousandeyes-operator.yaml](./operator.yaml)

3 Install the operator:
```
$ ./install-operator.sh
```

4. Verify installation status

Use command ```kubectl get pods``` to check the ThousandEyes Operator deploy status:
```
$ kubectl get pods
NAME                                          READY   STATUS    RESTARTS   AGE
thousandeyes-operator-564b5d75d-jllzk         1/1     Running   0          108s
```

### Deploy ThousandEyes CRDs

The configuration of ThousandEyes test setup should be described in ThousandEyes CRD. You will find all the manifests in [ThousandEyes CRDs](./config/crd/bases) folder.

1. To deploy the ThousandEyes CRDs on your Kubernetes cluster, please run the following script:

```
$ kubectl apply -f config/crd/bases/thousandeyes.devnet.cisco.com_pageloadtests.yaml
```

2.  Use command ```kubectl get crd``` to check the ThousandEyes CRD deploy status: 
```
$ kubectl get crd
NAME                                                CREATED AT
pageloadtests.thousandeyes.devnet.cisco.com         2021-06-17T05:36:22Z
```
### Run a page load test
1. To create a page load test,you need to deploy the ThousandEyes custom resource on your Kubernetes cluster, please run the following script:
```
$ kubectl apply -f config/samples/devnet_v1alpha1_pageloadtest.yaml
```
You will find the test with the basic settings configured in the custom resource in ThousandEyes dashboard.

2. To update the configuration of the page load test, just update the specific field of sepc of [custom resource](./config/samples) then deploy it on your Kubernetes cluster, please run the following script:
```
$ kubectl apply -f config/samples/devnet_v1alpha1_pageloadtest.yaml
```
3. To delete the page load test, please run the following script:
```
$ kubectl delete -f config/samples/devnet_v1alpha1_pageloadtest.yaml
```

To run a web transaction test, follow the steps above.

## Reference
1. [ThousandEyes Getting Started](https://docs.thousandeyes.com/product-documentation/getting-started)
2. [ThousandEyes Test MetaData](https://developer.thousandeyes.com/v6/tests/#/test_metadata)












