package main

/*
loadtest_thrift093/go$ GOPATH=`pwd` go build -v -o ./bin/server simple_thrift_service/cmd/server
*/

import (
	"flag"
	"fmt"
	"log"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"

	"simple_thrift_service/thrift_gen/hello"

	"simple_thrift_service/service"
)

const DIMENSION = 4

type HelloServiceHandler struct {
	docStore *service.DocStore
}

func (*HelloServiceHandler) SendMessage(request *hello.HelloRequest) (*hello.HelloResponse, error) {
	resp := hello.NewHelloResponse()
	resp.Message = request.Message
	log.Printf("%s\n", *request.Message)
	return resp, nil
}

func (this *HelloServiceHandler) GetRelevance(request *hello.HelloRequest) (*hello.HelloResponse, error) {
	resp := hello.NewHelloResponse()
	resp.Message = request.Message
	resp.Results = this.docStore.GetRelatedDocs(*request.InputID, 5)
	log.Printf("%s\n", *request.Message)
	return resp, nil
}

func NewHelloServiceHandler() *HelloServiceHandler {
	return &HelloServiceHandler{
		docStore: service.NewDocStore(DIMENSION),
	}
}

func RunServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	handler := NewHelloServiceHandler()
	processor := hello.NewHelloServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	return server.Serve()
}

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	port := flag.Int("port", 9394, "The port to listen to")
	flag.Parse()

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	addr := fmt.Sprintf("localhost:%d", int32(*port))
	log.Printf("Starting server in %s\n", addr)
	RunServer(transportFactory, protocolFactory, addr)
}
