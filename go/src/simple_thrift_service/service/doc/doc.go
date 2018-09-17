package doc

type DocVec struct {
	Vec []float64
}

// DotProd calculates dot product for vector v1 and v2.
func DotProd(v1, v2 *DocVec) float64 {
	score := 0.0
	for i := 0; i < len(v1.Vec); i++ {
		score += v1.Vec[i] * v2.Vec[i]
	}
	return score
}
