package handlers

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/dongdongjssy/order-service/model"

	log "github.com/sirupsen/logrus"
)

func benchmark(b *testing.B, size int) {
	// initialize a list of orders based on size
	orders := make([]model.Order, size)
	for i := 0; i < size; i++ {
		orders[i] = model.Order{
			CustomerId: fmt.Sprintf("customer%d", i),
			OrderId:    fmt.Sprintf("order%d", i),
			Timestamp:  fmt.Sprint(time.Now().Unix()),
			Items: []model.Item{
				{ItemId: fmt.Sprintf("item%d", i), CostEur: float64(i)},
			},
		}
	}

	// temporarily disable logging
	log.SetOutput(io.Discard)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transformOrders(&orders)
	}

	// resume logging
	log.SetOutput(os.Stdout)
}

func BenchmarkTransformOrders1(b *testing.B)      { benchmark(b, 1) }
func BenchmarkTransformOrders10(b *testing.B)     { benchmark(b, 10) }
func BenchmarkTransformOrders100(b *testing.B)    { benchmark(b, 100) }
func BenchmarkTransformOrders1000(b *testing.B)   { benchmark(b, 1000) }
func BenchmarkTransformOrders10000(b *testing.B)  { benchmark(b, 10000) }
func BenchmarkTransformOrders100000(b *testing.B) { benchmark(b, 100000) }
