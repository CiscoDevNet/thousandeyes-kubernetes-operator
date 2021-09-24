## Web Transaction Test (Using Annotations on Service)

In this example, Let`s deploy a **Service**, we will add annotations on it to run a **Web Transaction** test monitoring **Cisco DevNet homepage**.

Two options to run a test.

### Option 1: Run a Web Transaction test using `thousandeyes.devnet.cisco.com/test-url`

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
   labels:
      run: nginx
spec:
   selector:
      run: nginx
   ports:
      - port: 80
        protocol: TCP
```
All the other settings will use [default values](web_transaction_cr.md#the-test-settings-specified-in-spec-are-defined-below)

1. Create a Web Transactions test
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_default_settings.yaml
   ```
   The test will be created on dashboard.

2. Update the settings of the Web Transactions test

   Modify `thousandeyes.devnet.cisco.com/test-url` in [config/samples/annotations/service_webtransactions_default_settings.yaml](../config/samples/annotations/service_webtransactions_default_settings.yaml#L7) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_default_settings.yaml
   ```
   You will find the url have been updated.

3. Delete the Web Transactions test

   Set `thousandeyes.devnet.cisco.com/test-type` to `none` in [config/samples/annotations/service_webtransactions_removal_settings.yaml](../config/samples/annotations/service_webtransactions_removal_settings.yaml#L6) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_removal_settings.yaml
   ```
   The test will be removed from ThousandEyes dashboard.

### Option 2: Run a Page Load test using `thousandeyes.devnet.cisco.com/test-spec`

This annotation follows [Web Transaction CR Spec definition](./web_transaction_cr.md#the-test-settings-specified-in-spec-are-defined-below)

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
           "interval": 1800,
           "agents": [
             {"agentName":"Tokyo, Japan (Trial)"},
             {"agentName":"Singapore (Trial) - IPv6"}
           ],
           "alertRules": [
             {"ruleName":"Default Web Transaction 2.0 Alert Rule"}
           ]
         }
      thousandeyes.devnet.cisco.com/test-script: |
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
   labels:
      run: nginx
spec:
   selector:
      run: nginx
   ports:
      - port: 80
        protocol: TCP
```

1. Create a Web Transactions test
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_customized_settings.yaml
   ```
   The test will be created on dashboard.

2. Update the settings of the Web Transactions test

   Modify `thousandeyes.devnet.cisco.com/test-spec` in [config/samples/annotations/service_webtransactions_customized_settings.yaml](../config/samples/annotations/service_webtransactions_customized_settings.yaml#L7) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_customized_settings.yaml
   ```
   You will find the settings have been updated.

3. Delete the Web Transactions test

   Set `thousandeyes.devnet.cisco.com/test-type` to `none` in [config/samples/annotations/service_webtransactions_removal_settings.yaml](../config/samples/annotations/service_webtransactions_removal_settings.yaml#L6) and redeploy.
   ```
   kubectl apply -f config/samples/annotations/service_webtransactions_removal_settings.yaml
   ```
   The test will be removed from ThousandEyes dashboard.





