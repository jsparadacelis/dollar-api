# Dollar Api

Calls to https://currencyfreaks.com/ to get dollar convertions.

You need to add a ```.env``` file like that:

```
\app
    - api.go
    - controllers.go
    - http_client.go
.env
```
This ```.env``` file has this data:

```
API_DOLLAR_TOKEN=<YOUR_TOKEN>
API_DOLLAR_URL=<API_URL>
```
