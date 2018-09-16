/*
Thrift for simple hello service.

thrift --out src/$PKG/hellothrift --gen go:package_prefix=$PKG/hellothrift src/$PKG/hellothrift/hello.thrift
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
