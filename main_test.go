package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dongdongjssy/order-service/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestOrderTransformApiErrorCases(t *testing.T) {
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

	var tests = []struct {
		input         []model.Order
		expectedError string
	}{
		{ordersMissingCustomerId, "error in field 'customerId': required"},
		{ordersMissingOrderId, "error in field 'orderId': required"},
		{ordersMissingTimestamp, "error in field 'timestamp': required"},
		{ordersMissingItems, "error in field 'items': required"},
		{ordersWithEmptyItems, "error in field 'items': should contains at least 1 element"},
		{ordersWithInvalidCost, "error in field 'costEur': should be greater than or equal to 0"},
	}

	for _, test := range tests {
		// build an order
		orderJson, _ := json.Marshal(test.input)

		// send request
		rec := postRequest(server, strings.NewReader(string(orderJson)))

		// assertion
		var response model.Response
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(http.StatusBadRequest, response.Code)
		assert.Equal("failed to parse request body", response.Message)
		assert.Equal(test.expectedError, response.Errors[0])
	}

}

// send post request to transform order
func postRequest(server *gin.Engine, body *strings.Reader) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/orders/transform", body)
	server.ServeHTTP(rec, req)

	return rec
}
