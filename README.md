tinkoff-invest-status
=====================

Works with [Tinkoff Invest OpenAPI](https://github.com/TinkoffCreditSystems/invest-openapi). Returns portfolio and
currencies to STDOUT in JSON format.

How to use
----------

1) Obtain Tinkoff Invest token from [https://www.tinkoff.ru/invest/settings/](https://www.tinkoff.ru/invest/settings/)

2) Download the latest release

    ```shell
    cd $HOME
    wget https://github.com/a0s/tinkoff-invest-status/releases/latest/download/tinkoff-invest-status
    chmod +x tinkoff-invest-status
    ```

3) Run it and get full portfolio
    ```shell
    ./tinkoff-invest-status --token=%YOUR_TOKEN%
   
    # Select summary by USD with jq 
    ./tinkoff-invest-status --token=%YOUR_TOKEN% | jq '.[].Summary.USD'
    ```

Using with [MTMR](https://github.com/Toxblh/MTMR)
-------------------------------------------------

Add this into MTMR config

```json
{
  "type": "shellScriptTitledButton",
  "width": 150,
  "refreshInterval": 15,
  "source": {     
    "inline": "echo \"USD $($HOME/tinkoff-invest-status --token=%YOUR_TOKEN% | /usr/local/bin/jq '.[].Summary.USD')\""
  },
  "actions": [],
  "align": "right",
  "bordered": false
}
```

Known problems
-------------
 - Tinkoff Invest API has some limits on request frequency. In some cases, you need to increase `refreshInterval` up to 30s.
