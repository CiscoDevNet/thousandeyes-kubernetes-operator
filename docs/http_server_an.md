## HTTP Server Test (Using Annotations on Ingress)

In this example, Let`s deploy an **Ingress** used to expose **Nginx web app** to public.

We will add annotations on it to run a **HTTP Server** test monitoring **Cisco DevNet homepage**.

1. Create a Nginx web app
   ```
   kubectl apply -f config/samples/nginx.yaml
   ```
2. Check the Nginx pod status
   ```
    kubectl get pods -A | grep nginx
    default         nginx-6976ddb986-rxqv6                          1/1     Running     0          12s
   ```
3. Install Nginx Ingress Controller
   ```
   kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
   ```
4. Check the Ingress pod status
   ```
   kubectl get pods -A | grep ingress-nginx
     ingress-nginx   ingress-nginx-admission-create-2n62c        0/1     Completed   0          66s
     ingress-nginx   ingress-nginx-admission-patch-fwnlg         0/1     Completed   1          66s
     ingress-nginx   ingress-nginx-controller-68649d49b8-62zvc   1/1     Running     0          66s
   ```

Now we are ready to run a test.

Two options to run a test.

### Option 1: Run a HTTP Server test using `thousandeyes.devnet.cisco.com/test-url`

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
All the other settings will use [default values](http_server_cr.md#the-test-settings-specified-in-spec-are-defined-below)

1. Create a HTTP Server Test
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_default_settings.yaml
   ```
   The test will be created on dashboard.

2. Update the HTTP Server test url

   Modify `thousandeyes.devnet.cisco.com/test-url` in [config/samples/annotations/ingress_httpserver_default_settings.yaml](../config/samples/annotations/ingress_httpserver_default_settings.yaml#L10) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_default_settings.yaml 
   ```   
   You will find the url has been updated.

3. Delete the HTTP Server test

   Set `thousandeyes.devnet.cisco.com/test-type` to `none` in [config/samples/annotations/ingress_httpserver_removal.yaml](../config/samples/annotations/ingress_httpserver_removal.yaml#L8) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_removal.yaml
   ```
   The test will be removed from ThousandEyes dashboard.

### Option 2: Run a HTTP Server test using `thousandeyes.devnet.cisco.com/test-spec`

This annotation follows [HTTPServer CR Spec definition](./http_server_cr.md#the-test-settings-specified-in-spec-are-defined-below)

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
           "interval": 1800,
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
1. Create a HTTP Server test
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_customized_settings.yaml
   ```
   The test will be created on dashboard.

2. Update the settings of the HTTP Server test

   Modify `thousandeyes.devnet.cisco.com/test-spec` in [config/samples/annotations/ingress_httpserver_customized_settings.yaml](../config/samples/annotations/ingress_httpserver_customized_settings.yaml#L8) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_customized_settings.yaml
   ```
   You will find the settings have been updated.

3. Delete the HTTP Server test
   
   Just set `thousandeyes.devnet.cisco.com/test-type` to `none` in [config/samples/annotations/ingress_httpserver_removal.yaml](../config/samples/annotations/ingress_httpserver_removal.yaml#L8) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/ingress_httpserver_removal.yaml
   ```
   The test will be removed from ThousandEyes dashboard.


   









