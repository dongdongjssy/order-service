package handlers

import (
	"fmt"
	"net/http"
	"unicode"

	"github.com/dongdongjssy/order-service/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const (
	API_SUCCESS = "success"

	ERR_INVALID_BODY          = "failed to parse request body"
	ERR_INVALID_FIELD         = "error in filed '%s': %s"
	ERR_SHOULD_BE_GTE         = "should be greater than or equal to %s"
	ERR_MIN_NOT_SATISFIED     = "should contains at least %s element"
	ERR_DUPLICATED_ORDER      = "duplicated order found"
	ERR_INTERNAL_SERVER_ERROR = "internal server error"
)

// It transforms a list of orders to the format of a list of customer items which
// consists of item details, the count of purchased items, and total cost etc.
func TransformOrders(ctx *gin.Context) {
	// parse request body
	var orders []model.Order
	if err := ctx.ShouldBindJSON(&orders); err != nil {
		// return error details
		errors := parseValidationErrors(&err)
		if len(*errors) > 0 {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Code:    http.StatusBadRequest,
				Message: ERR_INVALID_BODY,
				Errors:  *errors,
			})
			return
		}

	}

	// transform orders
	if summaries, err := transformOrders(&orders); err != nil {

	} else {
		ctx.JSON(http.StatusOK, model.Response{
			Code:    http.StatusOK,
			Message: API_SUCCESS,
			Data:    *summaries,
		})
	}
}

// Parse validation errors with a list of readable strings
//
// The method abstracts failed fields and reasons from validation errors and create a list
// of human readable error messages. e.g.
// [
//
//	"error in field 'customerId': required",
//	"error in field 'costEur': should be greater or equal to 0"
//
// ]
func parseValidationErrors(err *error) *[]string {
	var errors []string
	if sliceValidationErrs, ok := (*err).(binding.SliceValidationError); ok {
		for _, sve := range sliceValidationErrs {
			if validationErrs, ok := sve.(validator.ValidationErrors); ok {
				for _, ve := range validationErrs {
					switch ve.Tag() {
					case "required":
						errors = append(errors,
							fmt.Sprintf(ERR_INVALID_FIELD, toLowerFirstChar(ve.Field()), ve.Tag()))
					case "gte":
						errors = append(errors,
							fmt.Sprintf(ERR_INVALID_FIELD, toLowerFirstChar(ve.Field()), fmt.Sprintf(ERR_SHOULD_BE_GTE, ve.Param())))
					case "min":
						errors = append(errors,
							fmt.Sprintf(ERR_INVALID_FIELD, toLowerFirstChar(ve.Field()), fmt.Sprintf(ERR_MIN_NOT_SATISFIED, ve.Param())))
					}
				}
			}
		}
	}
	return &errors
}

// Calculate the total purchased amount of a customer.
//
// This method provides a solution similar as reduce method for array in other programming languages.
func reduceAmount(array *[]model.Item, initial float64, f func(float64, *model.Item) float64) float64 {
	result := initial
	for _, v := range *array {
		result = f(result, &v)
	}
	return result
}

// Make first letter of a string as lowercase.
//
// Structure field names start with upper case due to json binding but in json field names usually
// start with lower case, so when validation error happens in any fields, we update the case of
// first letter to make the error messages consist with the json format from the request.
func toLowerFirstChar(str string) string {
	if len(str) == 0 {
		return str
	}

	strSlice := []rune(str)
	strSlice[0] = unicode.ToLower(strSlice[0])
	return string(strSlice)
}

func transformOrders(orders *[]model.Order) (*[]model.Summary, error) {
	// check duplication and aggregate customer orders
	// TODO: make it concurrent
	customerOrders := make(map[string]map[string]*model.Order)
	for _, order := range *orders {
		if customerOrders[order.CustomerId] != nil {
			if customerOrders[order.CustomerId][order.OrderId] != nil {
				continue
			} else {
				customerOrders[order.CustomerId][order.OrderId] = &order
			}
		} else {
			customerOrders[order.CustomerId] =
				map[string]*model.Order{order.OrderId: &order}
		}
	}

	// store a list of summaries for all customers
	summaries := []model.Summary{}
	for cId, oList := range customerOrders {
		amount := 0.0
		items := []model.Item{}
		for _, o := range oList {
			items = append(items, o.Items...)
			amount = reduceAmount(&o.Items, amount, func(acc float64, i *model.Item) float64 { return acc + i.CostEur })
		}

		summaries = append(summaries, model.Summary{
			CustomerId:          cId,
			NbrOfPurchasedItems: len(items),
			Items:               items,
			TotalAmountEur:      amount,
		})
	}

	return &summaries, nil
}
