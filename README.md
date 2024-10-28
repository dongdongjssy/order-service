# order-service

This module provides 1 api for orders transform:

| Method | Endpoint             | Description                                            |
| ------ | -------------------- | ------------------------------------------------------ |
| POST   | /v1/orders/transform | transform a list of orders to a list of customer items |

## Getting started

To run/test locally

1. Clone code from this repository `git clone https://github.com/dongdongjssy/order-service.git`
2. Make sure you have latest [go](https://go.dev/) installed
3. In root directory, run `go run .`
4. An example of invoking the api with curl:
    ```sh
    curl -X POST -d '[{"customerId": "01","orderId": "50","timestamp": "1637245070513","items": [{"itemId": "20201","costEur": 2.5}]}]' localhost:8080/v1/orders/transform --header "Content-Type:application/json"
    ```

## Document

While service is running, the api docs can be found here: [Swagger docs](http://localhost:8080/swagger/index.html#/), it contains details including endpoint url, parameters, and responses etc.

## Auto tests

To run all auto tests, in root folder run command:

```sh
go test ./... -v
```
