/*
Thrift for simple hello service.

loadtest_thrift093/go/src $ thrift --out . --gen go simple_thrift_service/thrift_def/hello.thrift
*/

namespace go simple_thrift_service.thrift_gen.hello

struct HelloRequest {
    1: optional string message;
}

struct HelloResponse {
    1: optional string message;
}

service HelloService {
    HelloResponse sendMessage(1: HelloRequest request);
}
