package doc

import (
	"math"
	"testing"
)

func TestDotProd(t *testing.T) {
	v1 := &DocVec{
		Vec: []float64{1.0, 0.9, 0.8},
	}
	v2 := &DocVec{
		Vec: []float64{0.8, 0.7, 0.6},
	}
	actualScore := DotProd(v1, v2)
	expectScore := 1.91
	if math.Abs(actualScore-expectScore) > 0.00001 {
		t.Errorf("Expected score %f, but got %f", expectScore, actualScore)
	}
}
