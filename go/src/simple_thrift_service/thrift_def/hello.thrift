/*
Thrift for simple hello service.

loadtest_thrift093/go/src $ thrift --out . --gen go simple_thrift_service/thrift_def/hello.thrift
*/

namespace go simple_thrift_service.thrift_gen.hello

struct HelloRequest {
    1: optional string message;
    2: optional i32 input_id;
    3: optional i16 dimension;
}

struct HelloResponse {
    1: optional string message;
    2: optional list<HelloDoc> results;
}

// doc info
struct HelloDoc {
	1: i32 doc_id;
	2: list<double> vec;
	3: optional double score;
}

service HelloService {
    // Simple message process and response
    HelloResponse sendMessage(1: HelloRequest request);
    // Get relevant content by given input_id
    HelloResponse getRelevance(1: HelloRequest request);
}
