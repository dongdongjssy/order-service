# order-service

The order service is used to transform orders. It accepts a list of customer orders and then transforms them into a list of objects, each object corresponding to an individual customer. The transformed object contains the number of items a customer purchased, the total cost amount of all items, and details of each purchased items etc.

1. It receives orders from other services
2. It calls auth service for authentication before transforming
3. It sends result to other services for further processing (e.g. payment service)

This module provides 1 POST api:

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

To run all auto tests(unit + integration), in root folder run command:

```sh
go test ./... -v

--- PASS: TestOrdersTransformAPISuccess (0.00s)
    --- PASS: TestOrdersTransformAPISuccess/duplicated_order (0.00s)
    --- PASS: TestOrdersTransformAPISuccess/aggregate_orders (0.00s)
    --- PASS: TestOrdersTransformAPISuccess/single_order (0.00s)
    --- PASS: TestOrdersTransformAPISuccess/multiple_orders (0.00s)
--- PASS: TestOrdersTransformAPIErrorCases (0.00s)
    --- PASS: TestOrdersTransformAPIErrorCases/invalid_cost_value (0.00s)
    --- PASS: TestOrdersTransformAPIErrorCases/missing_customer_id (0.00s)
    --- PASS: TestOrdersTransformAPIErrorCases/missing_order_id (0.00s)
    --- PASS: TestOrdersTransformAPIErrorCases/missing_timestamp (0.00s)
    --- PASS: TestOrdersTransformAPIErrorCases/missing_items (0.00s)
    --- PASS: TestOrdersTransformAPIErrorCases/empty_items (0.00s)
--- PASS: TestReduceAmount (0.00s)
    --- PASS: TestReduceAmount/calculate_total_amount_of_all_items (0.00s)
--- PASS: TestFirstCharToLowercase (0.00s)
    --- PASS: TestFirstCharToLowercase/change_first_letter_of_a_string_to_lowercase (0.00s)
    --- PASS: TestFirstCharToLowercase/change_an_empty_string (0.00s)
```

## Benchmark

To run benchmark tests, in handlers folder run command:

```sh
go test -bench=TransformOrders
```

Following is an example of time used to transform 1, 10, 100, 1000, 10000, 100000 orders

```
BenchmarkTransformOrders1-16               98226             13388 ns/op
BenchmarkTransformOrders10-16              62355             19745 ns/op
BenchmarkTransformOrders100-16             10000            100655 ns/op
BenchmarkTransformOrders1000-16             1143           1043607 ns/op
BenchmarkTransformOrders10000-16             105          11145297 ns/op
BenchmarkTransformOrders100000-16              8         125425575 ns/op
```
