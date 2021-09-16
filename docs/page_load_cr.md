## Page Load Test (Apply CR)

In this example, we will run a Page Load test to monitor **Cisco DevNet homepage**.

### 1. Define Custom Resource (CR)
You can define a Page Load CR below based on [Page Load CRD](../config/crd/bases/thousandeyes.devnet.cisco.com_pageloadtests.yaml) (CRD):

CR: [**config/samples/cr/devnet_v1alpha1_pageloadtest.yaml**](../config/samples/cr/devnet_v1alpha1_pageloadtest.yaml)
```yaml
apiVersion: thousandeyes.devnet.cisco.com/v1alpha1
kind: PageLoadTest
metadata:
   # the unique test name
   name: pageload-devnet-homepage
spec:
   url: https://developer.cisco.com/
   interval: 300
   httpInterval: 300
   agents:
      - agentName: Tokyo, Japan (Trial)
      - agentName: Singapore (Trial) - IPv6
   alertRules:
      - ruleName: Default Page Load Alert Rule
```
The test settings specified in **spec** are defined below:

| Field        | Test Creation| Test Update | Data Type | Default Values | Notes |
|--------------|--------------|-------------|----------|-----------------|-------|
|url           | Required     | n/a         | string   |                 | target for the test
|interval      | Required     |	Optional    | integer  | 120             | value in seconds. Accpeted Values:[120, 300, 600, 900, 1800, 3600]
|httpInterval  | Required     | Optional    | integer  | 120             | value in seconds.Accpeted Values:[120, 300, 600, 900, 1800, 3600].It can not be larger than the interval value; defaults to the same value as interval
|agents        | Required     | Optional    | array of agentName|        |
|agentName     | Required     | Optional    | string   | Tokyo, Japan (Trial), Singapore (Trial) - IPv6   | cloud agent name
|alertRules    | Optional     | Optional    | array of ruleName|         | if this field is not specified, The default alert rules will be used.
|ruleName      | Optional     | Optional    | string   |                 | alert rule name

For more details, please refer to **Page Load** in [ThousandEyes Test Metadata](https://developer.thousandeyes.com/v6/tests/#/test_metadata).

### 2. Run a Page Load Test

Run the commands respectively, check the test on ThoudandEyes dashboard.

1. Create a Page Load test
    ```
    kubectl apply -f config/samples/cr/devnet_v1alpha1_pageloadtest.yaml
    ```
   The test will be created on dashboard.

2. Update settings of this test

   Modify fields in **spec** in [devnet_v1alpha1_pageloadtest](../config/samples/cr/devnet_v1alpha1_pageloadtest.yaml) and redeploy.
    ```
    kubectl apply -f config/samples/cr/devnet_v1alpha1_pageloadtest.yaml
    ```
   You will find the settings have been updated.   

3. Delete this test
    ```
    kubectl delete -f config/samples/cr/devnet_v1alpha1_pageloadtest.yaml
    ```
  The test will be removed from ThousandEyes dashboard.






