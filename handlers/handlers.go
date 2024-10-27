package handlers

import (
	"net/http"

	"github.com/dongdongjssy/order-service/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	API_SUCCESS = "success"

	ERR_INVALID_BODY     = "failed to parse request body"
	ERR_DUPLICATED_ORDER = "duplicated order found"
)

func TransformOrders(ctx *gin.Context) {
	// parse request body
	var orders []model.Order
	if err := ctx.ShouldBindJSON(&orders); err != nil {
		// if find error, return parse error details
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)
			for _, e := range validationErrs {
				errors[e.Field()] = e.Tag()
			}
			ctx.JSON(http.StatusBadRequest, model.Response{
				Code:    http.StatusBadRequest,
				Message: ERR_INVALID_BODY,
				Error:   errors,
			})
			return
		}
	}

	// transform orders
	summaries := make([]model.Summary, len(orders))
	for index, order := range orders {
		for i := range order.Items {
			order.Items[i].CustomerId = order.CustomerId
		}

		summary := model.Summary{
			CustomerId:          order.CustomerId,
			NbrOfPurchasedItems: len(order.Items),
			Items:               order.Items,
			TotalAmountEur:      reduceAmount(order.Items, 0, func(acc float64, i model.Item) float64 { return acc + i.CostEur }),
		}

		summaries[index] = summary
	}

	ctx.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: API_SUCCESS,
		Data:    summaries,
	})
}

func reduceAmount(
	array []model.Item,
	initial float64,
	f func(float64, model.Item) float64,
) float64 {
	result := initial
	for _, v := range array {
		result = f(result, v)
	}
	return result
}
