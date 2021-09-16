## HTTP Server Test (Add Annotations to Ingress)

In this example, We will deploy an **Ingress** with annotations, it will run a **HTTP Server** test to monitor **Cisco DevNet homepage**.

This Ingress is used to expose **Nginx web app** to public. 

So First, we need to create an Nginx web app.

1. Create an Nginx web app
   ```
   cd thousandeyes-operator
   kubectl apply -f config/samples/nginx.yaml
   ```
2. Check the Nginx pod status
   ```
    kubectl get pods -A | grep nginx
    default         nginx-6976ddb986-rxqv6                          1/1     Running     0          12s
   ```

Second, we need to install an ingress controller on local.

### Install Ingress Controller Locally

There are multiple Ingress controllers, we will use Nginx Ingress Controller as an example.

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

Based on the pre-requisites, now we can run our test.

### Run a HTTP Server Test

1. If you want to customize the test settings, you can add **thousandeyes.devnet.cisco.com/test-spec** to Ingress.

Ingress: [**config/samples/annotations/ingress_httpserver_customized_settings.yaml**](../config/samples/annotations/ingress_httpserver_customized_settings.yaml)
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
   # the unique test name
   name: ingress-httpserver
   annotations:
      thousandeyes.devnet.cisco.com/test-type: http-server
      thousandeyes.devnet.cisco.com/test-spec: |
         {
           "url":"https://developer.cisco.com/",
           "interval": 300,
           "agents": [
              {"agentName":"Tokyo, Japan (Trial)"},
              {"agentName":"Singapore (Trial) - IPv6"}
           ],
           "alertRules": [
              {"ruleName":"Default HTTP Alert Rule"}
           ]
         }
spec:
   rules:
      - http:
           paths:
              - path: /
                pathType: Prefix
                backend:
                   service:
                      name: nginx-service
                      port:
                         number: 80
```
i. Create a HTTP Server test
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_customized_settings.yaml
   ```
   The test will be created on dashboard.

ii. Update the settings of the HTTP Server test

   Modify **thousandeyes.devnet.cisco.com/test-spec** in [Ingress resource](../config/samples/annotations/ingress_httpserver_customized_settings.yaml#L8) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_customized_settings.yaml
   ```
   You will find the settings have been updated.

iii. Delete the HTTP Server test
   
   Just set **thousandeyes.devnet.cisco.com/test-type** to **none** in [Ingress resource](../config/samples/annotations/ingress_httpserver_removal.yaml#L8) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_removal.yaml
   ```
   The test will be removed from ThousandEyes dashboard.

2. You can add **thousandeyes.devnet.cisco.com/test-url** only to Ingress.

   The test will be created with settings by default.

Ingress: [**config/samples/annotations/ingress_httpserver_default_settings.yaml**](../config/samples/annotations/ingress_httpserver_default_settings.yaml)

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
   # the unique test name
   name: ingress-httpserver
   annotations:
      thousandeyes.devnet.cisco.com/test-type: http-server
      thousandeyes.devnet.cisco.com/test-url: https://developer.cisco.com/
spec:
   rules:
      - http:
           paths:
              - path: /
                pathType: Prefix
                backend:
                   service:
                      name: nginx-service
                      port:
                         number: 80
```
i. Create a HTTP Server Test
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_default_settings.yaml
   ```
   The test will be created on dashboard.

ii. Update the **url** of the HTTP Server test

   Modify **thousandeyes.devnet.cisco.com/test-url** in [Ingress resource](../config/samples/annotations/ingress_httpserver_default_settings.yaml#L10) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_default_settings.yaml 
   ```   
   You will find the url have been updated.

iii. Delete the HTTP Server test
   
   Just set **thousandeyes.devnet.cisco.com/test-type** to **none** in [Ingress resource](../config/samples/annotations/ingress_httpserver_removal.yaml#L8) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_removal.yaml
   ```
   The test will be removed from ThousandEyes dashboard.

The usage of annotations applies to **Kubernetes Service** resource as well.

   









