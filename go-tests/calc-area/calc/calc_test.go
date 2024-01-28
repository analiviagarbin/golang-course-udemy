package calc

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retang", func(t *testing.T) {
		ret := Retang{ 10, 12 }
		expectedRes := float64(120)
		resCalc := ret.Area()

		if resCalc != expectedRes {
			t.Fatal("Invalid Area")
		}
	})

	t.Run("Circle", func(t *testing.T) {
		cir := Circle{ 10 }
		expectedRes := float64(math.Pi * 100)
		resCalc := cir.Area()

		if resCalc != expectedRes {
			t.Fatal("Invalid Area")
		}
	})
}