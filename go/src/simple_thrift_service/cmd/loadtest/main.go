package main

/*
loadtest_thrift093/go$ GOPATH=`pwd` go build -v -o ./bin/loadtest simple_thrift_service/cmd/loadtest
loadtest_thrift093/go$ ./bin/loadtest --num_request=20
*/

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/lenjoy/bender"
	"github.com/lenjoy/bender/hist"
	bthrift "github.com/lenjoy/bender/thrift"

	"simple_thrift_service/thrift_gen/hello"

	"simple_thrift_service/service"
	"simple_thrift_service/service/randgen"
)

func SyntheticRequests(query string, topK int16, n int) chan interface{} {
	c := make(chan interface{}, 100)
	go func() {
		for i := 0; i < n; i++ {
			request := hello.NewHelloRequest()
			inputMsg := fmt.Sprintf("%s %d", query, i)
			inputID := int32(i)
			reqTopK := int16(int(topK) + randgen.GetInt(5))
			request.Message = &inputMsg
			request.InputID = &inputID
			request.TopK = &reqTopK
			c <- request
		}
		close(c)
	}()
	return c
}

func HelloExecutor(request interface{}, transport thrift.TTransport) (interface{}, error) {
	pFac := thrift.NewTBinaryProtocolFactoryDefault()
	client := hello.NewHelloServiceClientFactory(transport, pFac)
	// return client.SendMessage(request.(*hello.HelloRequest))
	response, err := client.GetRelevance(request.(*hello.HelloRequest))
	service.PrintResp(response)
	return response, err
}

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	serverAddr := flag.String("server_addr", "localhost:9394", "The server address to talk to")
	query := flag.String("query", "hello, loadtest", "The input message")
	numRequest := flag.Int64("num_request", 10, "The num of requests sending to server")
	topK := flag.Int64("top_k", 3, "The num of results returning from server")
	flag.Parse()

	intervals := bender.ExponentialIntervalGenerator(10.0)
	requests := SyntheticRequests(*query, int16(*topK), int(*numRequest))
	timeout := 20 * time.Second
	exec := bthrift.NewThriftRequestExec(thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory()), HelloExecutor, timeout, *serverAddr)
	recorder := make(chan interface{}, 128)
	bender.LoadTestThroughput(intervals, requests, exec, recorder)
	l := log.New(os.Stdout, "", log.LstdFlags)
	h := hist.NewHistogram(60000, 1000000)
	bender.Record(recorder, bender.NewLoggingRecorder(l), bender.NewHistogramRecorder(h))
	log.Println(h)
}
