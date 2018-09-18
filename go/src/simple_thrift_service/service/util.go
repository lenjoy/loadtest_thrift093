package service

import (
	"fmt"
	"log"

	"simple_thrift_service/thrift_gen/hello"
)

func PrintResp(response *hello.HelloResponse) {
	log.Println(fmt.Sprintf("Message: [%s]", *response.Message))
	for _, result := range response.Results {
		log.Println(fmt.Sprintf("doc_id=%d\tscore=%f, %v", result.DocID, *result.Score, result.Vec))
	}
}
