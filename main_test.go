package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dongdongjssy/order-service/handlers"
	"github.com/dongdongjssy/order-service/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestOrdersTransformAPISuccess(t *testing.T) {
	assert := assert.New(t)
	server := setupRouter()

	orders := []model.Order{
		{CustomerId: "01", OrderId: "10", Timestamp: "1637245070513", Items: []model.Item{
			{ItemId: "20201", CostEur: 2},
			{ItemId: "20202", CostEur: 4},
		}},
		{CustomerId: "02", OrderId: "20", Timestamp: "1637245070513", Items: []model.Item{
			{ItemId: "20203", CostEur: 2.1},
		}},
	}

	t.Run("transform success", func(t *testing.T) {
		// build an order
		ordersJson, _ := json.Marshal(orders)

		// send request
		rec := postRequest(server, strings.NewReader(string(ordersJson)))

		// assertion
		var response model.Response
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(http.StatusOK, response.Code)
		assert.Equal(handlers.API_SUCCESS, response.Message)
		assert.Equal(2, len(response.Data))

		for _, data := range response.Data {
			if data.CustomerId != orders[0].CustomerId && data.CustomerId != orders[1].CustomerId {
				t.Errorf(
					"invalid response, expected customer id with value of either '%s' or '%s'",
					orders[0].CustomerId,
					orders[1].CustomerId,
				)
			}

			if data.CustomerId == orders[0].CustomerId {
				assert.Equal(2, data.NbrOfPurchasedItems)
				assert.Equal(float64(6), data.TotalAmountEur)
				assert.Equal(2, len(data.Items))
			}

			if data.CustomerId == orders[1].CustomerId {
				assert.Equal(1, data.NbrOfPurchasedItems)
				assert.Equal(2.1, data.TotalAmountEur)
				assert.Equal(1, len(data.Items))
			}
		}
	})
}

func TestOrdersTransformAPIErrorCases(t *testing.T) {
	assert := assert.New(t)
	server := setupRouter()

	ordersMissingCustomerId := []model.Order{
		{OrderId: "50", Timestamp: "1637245070513", Items: []model.Item{
			{ItemId: "20201", CostEur: 2},
		}},
	}

	ordersMissingOrderId := []model.Order{
		{CustomerId: "01", Timestamp: "1637245070513", Items: []model.Item{
			{ItemId: "20201", CostEur: 2},
		}},
	}

	ordersMissingTimestamp := []model.Order{
		{CustomerId: "01", OrderId: "50", Items: []model.Item{
			{ItemId: "20201", CostEur: 2},
		}},
	}

	ordersMissingItems := []model.Order{{CustomerId: "01", OrderId: "50", Timestamp: "1637245070513"}}

	ordersWithEmptyItems := []model.Order{
		{CustomerId: "01", OrderId: "50", Timestamp: "1637245070513", Items: []model.Item{}},
	}

	ordersWithInvalidCost := []model.Order{
		{CustomerId: "01", OrderId: "50", Timestamp: "1637245070513", Items: []model.Item{
			{ItemId: "20201", CostEur: -2},
		}},
	}

	var tests = map[string]struct {
		input         []model.Order
		expectedError string
	}{
		"missing customer id": {ordersMissingCustomerId, "error in field 'customerId': required"},
		"missing order id":    {ordersMissingOrderId, "error in field 'orderId': required"},
		"missing timestamp":   {ordersMissingTimestamp, "error in field 'timestamp': required"},
		"missing items":       {ordersMissingItems, "error in field 'items': required"},
		"empty items":         {ordersWithEmptyItems, "error in field 'items': should contains at least 1 element"},
		"invalid cost value":  {ordersWithInvalidCost, "error in field 'costEur': should be greater than or equal to 0"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// build an order
			ordersJson, _ := json.Marshal(test.input)

			// send request
			rec := postRequest(server, strings.NewReader(string(ordersJson)))

			// assertion
			var response model.Response
			json.Unmarshal(rec.Body.Bytes(), &response)

			assert.Equal(http.StatusBadRequest, response.Code)
			assert.Equal(handlers.ERR_INVALID_BODY, response.Message)
			assert.Equal(test.expectedError, response.Errors[0])
		})
	}

}

// send post request to orders transform endpoint
func postRequest(server *gin.Engine, body *strings.Reader) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, ENDPOINT_ORDERS_TRANSFORM, body)
	server.ServeHTTP(rec, req)

	return rec
}
