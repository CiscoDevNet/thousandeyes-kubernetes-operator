## HTTP Server Test (Apply CR)

In this example, we will run a HTTP Server test to monitor **Cisco DevNet homepage**.

### 1. Define Custom Resource (CR)
We can define a HTTP Server CR below based on [HTTP Server CRD](../config/crd/bases/thousandeyes.devnet.cisco.com_httpservertests.yaml):

CR: [**config/samples/cr/devnet_v1alpha1_httpservertest.yaml**](../config/samples/cr/devnet_v1alpha1_httpservertest.yaml)
```yaml
apiVersion: thousandeyes.devnet.cisco.com/v1alpha1
kind: HTTPServerTest
metadata:
  # the test name
  name: httpserver-devnet-homepage
spec:
  url: https://developer.cisco.com/
  interval: 300
  agents:
    - agentName: Tokyo, Japan (Trial)
    - agentName: Singapore (Trial) - IPv6
  alertRules:
    - ruleName: Default HTTP Alert Rule
```
The test settings specified in **spec** are defined below:

| Field        | Test Creation| Test Update | Data Type | Default Values| Notes |
|--------------|--------------|-------------|-----------|---------------|-------|
|url           | Required     | n/a         | string    |               | target for the test
|interval      | Required     |	Optional    | integer   | 120           | value in seconds. Accpeted Values:[120, 300, 600, 900, 1800, 3600]
|agents        | Required     | Optional    | array of agentName|       |
|agentName     | Required     | Optional    | string    |Tokyo, Japan (Trial), Singapore (Trial) - IPv6 | cloud agent name
|alertRules    | Optional     | Optional    | array of ruleName|        | if this field is not specified, The default alert rules will be used.
|ruleName      | Optional     | Optional    | string    |               | alert rule name

For more details, please refer to **HTTP Server** in [ThousandEyes Test Metadata](https://developer.thousandeyes.com/v6/tests/#/test_metadata).

### 2. Run a HTTP Server Test

Run the commands respectively, check the test on ThoudandEyes dashboard.

1. Create a HTTP Server test
    ```
    kubectl apply -f config/samples/cr/devnet_v1alpha1_httpservertest.yaml
    ```
    The test will be created on dashboard.

2. Update settings of this test

   Modify fields in **spec** in [devnet_v1alpha1_httpservertest.yaml](../config/samples/cr/devnet_v1alpha1_httpservertest.yaml#L7) and redeploy.
    ```
    kubectl apply -f config/samples/cr/devnet_v1alpha1_httpservertest.yaml
    ```
   You will find the settings have been updated.

3. Delete this test
    ```
    kubectl delete -f config/samples/cr/devnet_v1alpha1_httpservertest.yaml
    ```
   The test will be removed from ThousandEyes dashboard.






