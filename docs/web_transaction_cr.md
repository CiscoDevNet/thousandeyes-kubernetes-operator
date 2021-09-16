## Web Transaction Test (Apply CR)

In this example, we will run a Web Transaction test to interact with **Cisco DevNet homepage**.

### 1. Define Custom Resource (CR)
You can define a Web Transaction CR below based on [Web Transaction CRD](../config/crd/bases/thousandeyes.devnet.cisco.com_webtransactiontests.yaml) (CRD):

CR: [**config/samples/cr/devnet_v1alpha1_webtransactiontest.yaml**](../config/samples/cr/devnet_v1alpha1_webtransactiontest.yaml)
```yaml
apiVersion: thousandeyes.devnet.cisco.com/v1alpha1
kind: WebTransactionTest
metadata:
  # the test name
  name: transactions-devnet-homepage
spec:
  url: https://developer.cisco.com/
  interval: 120
  agents:
    - agentName: Hong Kong (Trial)
    - agentName: Singapore (Trial) - IPv6
  alertRules:
    - ruleName: Default HTTP Alert Rule
  transactionScript: |
     import { By, Key } from 'selenium-webdriver';
     import { driver, test } from 'thousandeyes';
     runScript();
     async function runScript() {
         await configureDriver();
         const settings = test.getSettings();
         // Load page
         await driver.get(settings.url);
         await click(By.id(`offer-getstarted`));
     }
     async function configureDriver() {
       await driver.manage().setTimeouts({
           implicit: 7 * 1000, // If an element is not found, reattempt for this many milliseconds
       });
     }
      ... ...
```
The test settings specified in **spec** are defined below:

| Field        | Test Creation| Test Update | Data Type | Default Values | Notes
|--------------|--------------|-------------|-----------|----------------|-------|
|url           | Required     | n/a         | string    |                | target for the test
|interval      | Required     |	Optional    | integer   | 120            | value in seconds. Accpeted Values:[120, 300, 600, 900, 1800, 3600]
|agents        | Required     | Optional    | array of agentName|        |
|agentName     | Required     | Optional    | string    | Tokyo, Japan (Trial), Singapore (Trial) - IPv6 | cloud agent name
|alertRules    | Optional     | Optional    | array of ruleName|         | if this field is not specified, The default alert rules will be used.
|ruleName      | Optional     | Optional    | string    |                | alert rule name
|transactionScript|Required   | Optional    | string    |                | javaScript of a web transaction test.It could be generated via ThousandEyes Recorder. 

For more details, please refer to **Web Transaction** [ThousandEyes Test Metadata](https://developer.thousandeyes.com/v6/tests/#/test_metadata).

### 2. Run a Web Transaction Test

Run the commands respectively, check the test on ThoudandEyes dashboard.

1. Create a Web Transaction test
   ```
   kubectl apply -f config/samples/cr/devnet_v1alpha1_webtransactiontest.yaml
   ```
   The test will be created on dashboard.

2. Update settings of this test

   Modify fields in **spec** in [devnet_v1alpha1_webtransactiontest.yaml](../config/samples/cr/devnet_v1alpha1_webtransactiontest.yaml) and redeploy
    ```
    kubectl apply -f config/samples/cr/devnet_v1alpha1_webtransactiontest.yaml
    ```
    You will find the settings have been updated.

3. Delete this test
    ```
    kubectl delete -f config/samples/cr/devnet_v1alpha1_webtransactiontest.yaml
    ```
   The test will be removed from ThousandEyes dashboard.



