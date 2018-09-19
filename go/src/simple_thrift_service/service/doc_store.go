package service

import (
	//	"log"
	"sort"
	"sync"
	"time"

	"simple_thrift_service/thrift_gen/hello"

	"simple_thrift_service/service/doc"
	"simple_thrift_service/service/randgen"
)

type DocStore struct {
	// The dimension of the doc vector.
	Dim int16
}

func NewDocStore(dim int16) *DocStore {
	return &DocStore{
		Dim: dim,
	}
}

// GetVector gets vector by given docID.
func (this *DocStore) GetVector(docID int32) *doc.DocVec {
	// log.Printf("Begin GetVector doc_id[%d]", docID)
	// simulate the latency between [10, 20)
	cnt := time.Duration(10 + randgen.GetInt(10))
	time.Sleep(time.Millisecond * cnt)
	return randgen.GenVec(int(this.Dim))
}

// GetRelatedDocs gets related docs by given docID.
func (this *DocStore) GetRelatedDocs(docID int32, n int) []*hello.HelloDoc {
	seedDocVec := this.GetVector(docID)

	relatedDocIDs := randgen.GenArray(n + n) // get more candidates
	totalNum := len(relatedDocIDs)

	var wg sync.WaitGroup
	wg.Add(totalNum)
	docArrCh := make(chan *hello.HelloDoc, 100)
	for i := 0; i < totalNum; i++ {
		relDocID := relatedDocIDs[i]
		go func() {
			docVec := this.GetVector(relDocID)
			score := doc.DotProd(seedDocVec, docVec)
			doc := hello.NewHelloDoc()
			doc.DocID = relDocID
			doc.Vec = docVec.Vec
			doc.Score = &score
			docArrCh <- doc
			wg.Done()
			//			log.Printf("%d - doc_id[%d] done!", totalNum, relDocID)
		}()
	}
	wg.Wait()
	close(docArrCh)

	arr := []*hello.HelloDoc{}
	for doc := range docArrCh {
		arr = append(arr, doc)
	}

	sort.Slice(arr, func(i, j int) bool {
		// from higher to lower
		return *arr[i].Score > *arr[j].Score
	})

	return arr[:n]
}
