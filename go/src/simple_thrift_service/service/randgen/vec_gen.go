package randgen

import (
	"math/rand"
	"time"

	"simple_thrift_service/service/doc"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

func GenArray(n int) []int32 {
	ret := make([]int32, n)
	for i := 0; i < n; i++ {
		ret[i] = r.Int31()
	}
	return ret
}

func GenVec(n int) *doc.DocVec {
	ret := make([]float64, n)
	for i := 0; i < n; i++ {
		ret[i] = r.Float64()
	}

	return &doc.DocVec{
		Vec: ret,
	}
}

func GetInt(n int) int {
	return r.Intn(n)
}
