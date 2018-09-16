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

func Send(client *hello.HelloServiceClient, inputMessage string) error {
	request := hello.NewHelloRequest()
	request.Message = &inputMessage
	response, err := client.SendMessage(request)
	if err != nil {
		log.Printf("%s - %v\n", inputMessage, err)
		return err
	}
	log.Println(*response.Message)
	return nil
}

func RunClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, inputMessage string) (*hello.HelloServiceClient, error) {
	socket, err := thrift.NewTSocket(addr)
	if err != nil {
		return nil, err
	}

	transport := transportFactory.GetTransport(socket)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return nil, err
	}

	client := hello.NewHelloServiceClientFactory(transport, protocolFactory)
	err = Send(client, inputMessage)
	if err != nil {
		Check(err)
	}

	// TODO: figure out why returned `client` always suffers "Connection not open".
	return client, nil
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

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	addr := *serverAddr
	log.Printf("Talking to server at [%s]\n", addr)

	inputMessage := *query
	_, err := RunClient(transportFactory, protocolFactory, addr, inputMessage)
	if err != nil {
		log.Println("can not get client")
		Check(err)
	}
}
