## Web Transaction Test (Add Annotations to Service)

In this example, we will deploy a following **Service** with annotations, it will run a **Web Transaction** test to interact with **Cisco DevNet homepage**.

1. If we want to customize the test settings, we can add **thousandeyes.devnet.cisco.com/test-spec** to Service.

Service: [**config/samples/annotations/service_webtransactions_customized_settings.yaml**](../config/samples/annotations/service_webtransactions_customized_settings.yaml)

```yaml
apiVersion: v1
kind: Service
metadata:
   name: nginx-service-webtransactions
   annotations:
      thousandeyes.devnet.cisco.com/test-type: web-transactions
      thousandeyes.devnet.cisco.com/test-spec: |
         {
           "url":"https://developer.cisco.com/",
           "interval": 300,
           "agents": [
             {"agentName":"Tokyo, Japan (Trial)"},
             {"agentName":"Singapore (Trial) - IPv6"}
           ],
           "alertRules": [
             {"ruleName":"Default Web Transaction 2.0 Alert Rule"}
           ]
         }
      thousandeyes.devnet.cisco.com/test-script: |
         test
   labels:
      run: nginx
spec:
   selector:
      run: nginx
   ports:
      - port: 80
        protocol: TCP
```

i. Create a Web Transactions test
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_customized_settings.yaml
   ```
The test will be created on dashboard.

ii. Update the settings of the Web Transactions test

Modify **thousandeyes.devnet.cisco.com/test-spec** in [Service resource](../config/samples/annotations/service_webtransactions_customized_settings.yaml#L7) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_customized_settings.yaml
   ```
You will find the settings have been updated.

iii. Delete the Web Transactions test

Just set **thousandeyes.devnet.cisco.com/test-type** to **none** in [Service resource](../config/samples/annotations/service_webtransactions_removal_settings.yaml#L6) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_removal_settings.yaml
   ```
The test will be removed from ThousandEyes dashboard.

2. You can add **thousandeyes.devnet.cisco.com/test-url** only to Service.

   The test will be created with settings by default.

Service: [**config/samples/annotations/service_webtransactions_default_settings.yaml**](../config/samples/annotations/service_webtransactions_default_settings.yaml)

```yaml
apiVersion: v1
kind: Service
metadata:
   name: nginx-service-webtransactions
   annotations:
      thousandeyes.devnet.cisco.com/test-type: web-transactions
      thousandeyes.devnet.cisco.com/test-url: https://developer.cisco.com/
      thousandeyes.devnet.cisco.com/test-script: |
         test
   labels:
      run: nginx
spec:
   selector:
      run: nginx
   ports:
      - port: 80
        protocol: TCP
```

i. Create a Web Transactions test
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_default_settings.yaml
   ```
The test will be created on dashboard.

ii. Update the settings of the Web Transactions test

Modify **thousandeyes.devnet.cisco.com/test-url** in [Service resource](../config/samples/annotations/service_webtransactions_default_settings.yaml#L7) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_default_settings.yaml
   ```
You will find the url have been updated.

iii. Delete the Web Transactions test

Just set **thousandeyes.devnet.cisco.com/test-type** to **none** in [Service resource](../config/samples/annotations/service_webtransactions_removal_settings.yaml#L6) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_removal_settings.yaml
   ```
The test will be removed from ThousandEyes dashboard.

The usage of annotations applies to **Kubernetes Ingress** resource as well.



