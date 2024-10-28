package handlers

import (
	"testing"

	"github.com/dongdongjssy/order-service/model"
	"github.com/stretchr/testify/assert"
)

func TestReduceAmount(t *testing.T) {
	t.Run("calculate total amount of all items", func(t *testing.T) {
		items := []model.Item{
			{ItemId: "1", CostEur: 1},
			{ItemId: "2", CostEur: 2},
			{ItemId: "3", CostEur: 3},
		}

		total := reduceAmount(&items, 0, func(acc float64, i *model.Item) float64 { return acc + i.CostEur })
		assert.Equal(t, float64(6), total)
	})
}

func TestFirstCharToLowercase(t *testing.T) {
	t.Run("change first letter of a string to lowercase", func(t *testing.T) {
		result := toLowerFirstChar("CustomerId")
		assert.Equal(t, "customerId", result)
	})

	t.Run("change an empty string", func(t *testing.T) {
		result := toLowerFirstChar("")
		assert.Equal(t, "", result)
	})
}
