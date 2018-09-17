package service

import (
	"sort"

	"simple_thrift_service/thrift_gen/hello"

	"simple_thrift_service/service/doc"
	"simple_thrift_service/service/randgen"
)

const DIMENSION = 4

type DocStore struct {
	// The dimension of the doc vector.
	Dim int16
}

func NewDocStore() *DocStore {
	return &DocStore{
		Dim: DIMENSION,
	}
}

// GetVector gets vector by given docID.
func (this *DocStore) GetVector(docID int32) *doc.DocVec {
	return randgen.GenVec(int(this.Dim))
}

// GetRelatedDocs gets related docs by given docID.
func (this *DocStore) GetRelatedDocs(docID int32, n int) []*hello.HelloDoc {
	seedDocVec := this.GetVector(docID)

	arr := []*hello.HelloDoc{}
	relatedDocIDs := randgen.GenArray(n + n) // get more candidates
	for _, relDocID := range relatedDocIDs {
		docVec := this.GetVector(relDocID)
		score := doc.DotProd(seedDocVec, docVec)
		doc := hello.NewHelloDoc()
		doc.DocID = relDocID
		doc.Vec = docVec.Vec
		doc.Score = &score
		arr = append(arr, doc)
	}

	sort.Slice(arr, func(i, j int) bool {
		// from higher to lower
		return *arr[i].Score > *arr[j].Score
	})

	return arr[:n]
}
