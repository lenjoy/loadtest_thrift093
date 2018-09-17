package main

/*
loadtest_thrift093/go$ GOPATH=`pwd` go build -v -o ./bin/client simple_thrift_service/cmd/client
*/

import (
	"flag"
	"fmt"
	"log"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"

	"simple_thrift_service/thrift_gen/hello"
)

func Connect(transportFactory thrift.TTransportFactory, addr string) (thrift.TTransport, error) {
	log.Printf("Talking to server at [%s]\n", addr)
	socket, err := thrift.NewTSocket(addr)
	if err != nil {
		return nil, err
	}

	transport := transportFactory.GetTransport(socket)
	return transport, nil
}

func Send(protocolFactory thrift.TProtocolFactory, transport thrift.TTransport, inputMessage string) error {
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}

	client := hello.NewHelloServiceClientFactory(transport, protocolFactory)

	// SendMessage
	request := hello.NewHelloRequest()
	request.Message = &inputMessage
	response, err := client.SendMessage(request)
	if err != nil {
		log.Printf("%s - %v\n", inputMessage, err)
		return err
	}
	log.Println(*response.Message)

	// GetRelevance
	var dim int16 = 4
	var inputID int32 = 1001
	request.Dimension = &dim
	request.InputID = &inputID
	response, err = client.GetRelevance(request)
	if err != nil {
		log.Printf("%s - %v\n", inputMessage, err)
		return err
	}
	log.Println(response)
	return nil
}

func Check(err error) {
	panic(err)
}

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	serverAddr := flag.String("server_addr", "localhost:9394", "The server address to talk to")
	query := flag.String("query", "hello, world!", "The input message")
	flag.Parse()

	addr := *serverAddr
	inputMessage := *query

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, err := Connect(transportFactory, addr)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	err = Send(protocolFactory, transport, inputMessage)
	if err != nil {
		log.Println("can not get client")
		Check(err)
	}
}
