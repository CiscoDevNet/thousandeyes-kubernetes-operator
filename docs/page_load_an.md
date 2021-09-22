## Page Load Test (Add Annotations to Service)

In this example, we will deploy the following **Service** with annotations, it will run a **Page Load** test to monitor **Cisco DevNet homepage**.

1.If you want to customize the test settings, you can add **thousandeyes.devnet.cisco.com/test-spec** to Service.

This annotation follows [PageLoad CR Spec definition](./page_load_cr.md#the-test-settings-specified-in-spec-are-defined-below)

Service: [**config/samples/annotations/service_pageload_customized_settings.yaml**](../config/samples/annotations/service_pageload_customized_settings.yaml)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-service-pageload
  annotations:
    thousandeyes.devnet.cisco.com/test-type: page-load
    thousandeyes.devnet.cisco.com/test-spec: |
      {
        "url":"https://developer.cisco.com/",
        "interval": 1800,
        "httpInterval": 1800,
        "agents": [
           {"agentName":"Tokyo, Japan (Trial)"},
           {"agentName":"Singapore (Trial) - IPv6"}
        ],
        "alertRules": [
           {"ruleName":"Default Page Load Alert Rule"}
        ]
      }
  labels:
    run: nginx
spec:
  selector:
    run: nginx
  ports:
    - port: 80
      protocol: TCP
```

i. Create a Page Load test
   ```
   kubectl apply -f config/samples/annotations/service_pageload_customized_settings.yaml
   ```
   The test will be created on dashboard.

ii. Update the settings of the Page Load test

   Modify **thousandeyes.devnet.cisco.com/test-spec** in [Service resource](../config/samples/annotations/service_pageload_customized_settings.yaml#L7) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_pageload_customized_settings.yaml
   ```
   You will find the settings have been updated.

iii. Delete the Page Load test

   Just set **thousandeyes.devnet.cisco.com/test-type** to **none** in [Service resource](../config/samples/annotations/service_pageload_removal_settings.yaml#L6) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_pageload_removal_settings.yaml
   ```
   The test will be removed from ThousandEyes dashboard.

2.If you want to use the [default settings](page_load_cr.md#the-test-settings-specified-in-spec-are-defined-below), you can just add **thousandeyes.devnet.cisco.com/test-url** to Service.

Service: [**config/samples/annotations/service_pageload_default_settings.yaml**](../config/samples/annotations/service_pageload_default_settings.yaml)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-service-pageload
  annotations:
    thousandeyes.devnet.cisco.com/test-type: page-load
    thousandeyes.devnet.cisco.com/test-url: https://developer.cisco.com/
  labels:
    run: nginx
spec:
  selector:
    run: nginx
  ports:
    - port: 80
      protocol: TCP
```

i. Create a Page Load test
   ```
   kubectl apply -f config/samples/annotations/service_pageload_default_settings.yaml
   ```
   The test will be created on dashboard.

ii. Update the settings of the Page Load test

   Modify **thousandeyes.devnet.cisco.com/test-url** in [Service resource](../config/samples/annotations/service_pageload_default_settings.yaml#L7) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_pageload_default_settings.yaml
   ```
   You will find the url have been updated.

iii. Delete the Page Load test

   Just set **thousandeyes.devnet.cisco.com/test-type** to **none** in [Service resource](../config/samples/annotations/service_pageload_removal_settings.yaml#L6) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_pageload_removal_settings.yaml
   ```
The test will be removed from ThousandEyes dashboard.

The usage of annotations applies to **Kubernetes Ingress** resource as well.








