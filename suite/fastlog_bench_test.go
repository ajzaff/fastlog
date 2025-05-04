package suite

import (
	"math"
	"testing"

	"github.com/ajzaff/fastlog"
)

func BenchmarkFastLog2(b *testing.B) {
	for b.Loop() {
		for _, example := range suite {
			fastlog.Log2(example.X)
		}
	}
}

func BenchmarkMathLog2(b *testing.B) {
	for b.Loop() {
		for _, example := range suite {
			math.Log2(float64(example.X))
		}
	}
}
