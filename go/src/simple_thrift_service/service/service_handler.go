package service

import (
	"log"
	"time"

	metrics "github.com/rcrowley/go-metrics"

	"simple_thrift_service/thrift_gen/hello"
)

const DIMENSION = 4

var (
	numGetRelevance     metrics.Meter
	latencyGetRelevance metrics.Timer
)

func init() {
	numGetRelevance = metrics.NewMeter()
	metrics.Register("api.num.GetRelevance", numGetRelevance)
	latencyGetRelevance = metrics.NewTimer()
	metrics.Register("api.latency.GetRelevance", latencyGetRelevance)
}

type HelloServiceHandler struct {
	docStore *DocStore
}

func (*HelloServiceHandler) SendMessage(request *hello.HelloRequest) (*hello.HelloResponse, error) {
	resp := hello.NewHelloResponse()
	resp.Message = request.Message
	log.Printf("SendMessage: %s\n", *request.Message)
	return resp, nil
}

func (this *HelloServiceHandler) GetRelevance(request *hello.HelloRequest) (*hello.HelloResponse, error) {
	numGetRelevance.Mark(1)

	start := time.Now().UnixNano()
	log.Printf(" Begin GetRelevance: top[%d] %s\n", *request.TopK, *request.Message)

	resp := hello.NewHelloResponse()
	resp.Message = request.Message
	resp.Results = this.docStore.GetRelatedDocs(*request.InputID, int(*request.TopK))
	log.Printf("  End  GetRelevance: top[%d] %s\n", *request.TopK, *request.Message)

	elapsedInNano := time.Duration(time.Now().UnixNano() - start)
	latencyGetRelevance.Update(elapsedInNano)

	return resp, nil
}

func NewHelloServiceHandler() *HelloServiceHandler {
	return &HelloServiceHandler{
		docStore: NewDocStore(DIMENSION),
	}
}
