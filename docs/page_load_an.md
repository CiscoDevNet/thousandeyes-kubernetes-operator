## Page Load Test (Using Annotations on Service)

In this example, Let`s deploy a **Service**, we will add annotations on it to run a **Page Load** test monitoring **Cisco DevNet homepage**.

Two options to run a Page Load test.

### Option 1: Using `thousandeyes.devnet.cisco.com/test-url`

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
All the other settings will use [default values](page_load_cr.md#the-test-settings-specified-in-spec-are-defined-below)

1. Create a Page Load test
   ```
   kubectl apply -f config/samples/annotations/service_pageload_default_settings.yaml
   ```
   The test will be created on dashboard.

2. Update the settings of the Page Load test

   Modify `thousandeyes.devnet.cisco.com/test-url` in [config/samples/annotations/service_pageload_default_settings.yaml](../config/samples/annotations/service_pageload_default_settings.yaml#L7) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_pageload_default_settings.yaml
   ```
   You will find the url have been updated.

3. Delete the Page Load test

   Set `thousandeyes.devnet.cisco.com/test-type` to `none` in [config/samples/annotations/service_pageload_removal_settings.yaml](../config/samples/annotations/service_pageload_removal_settings.yaml#L6) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_pageload_removal_settings.yaml
   ```
   The test will be removed from ThousandEyes dashboard.

### Option 2: Using `thousandeyes.devnet.cisco.com/test-spec`

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

1. Create a Page Load test
   ```
   kubectl apply -f config/samples/annotations/service_pageload_customized_settings.yaml
   ```
   The test will be created on dashboard.

2. Update the settings of the Page Load test

   Modify `thousandeyes.devnet.cisco.com/test-spec` in [config/samples/annotations/service_pageload_customized_settings.yaml](../config/samples/annotations/service_pageload_customized_settings.yaml#L7) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_pageload_customized_settings.yaml
   ```
   You will find the settings have been updated.

3. Delete the Page Load test

   Set `thousandeyes.devnet.cisco.com/test-type` to `none` in [config/samples/annotations/service_pageload_removal_settings.yaml](../config/samples/annotations/service_pageload_removal_settings.yaml#L6) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_pageload_removal_settings.yaml
   ```
   The test will be removed from ThousandEyes dashboard.











