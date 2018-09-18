package service

import (
	"log"

	"simple_thrift_service/thrift_gen/hello"
)

const DIMENSION = 4

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
	resp := hello.NewHelloResponse()
	resp.Message = request.Message
	resp.Results = this.docStore.GetRelatedDocs(*request.InputID, int(*request.TopK))
	log.Printf("GetRelevance: %s\n", *request.Message)
	return resp, nil
}

func NewHelloServiceHandler() *HelloServiceHandler {
	return &HelloServiceHandler{
		docStore: NewDocStore(DIMENSION),
	}
}
