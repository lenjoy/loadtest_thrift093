
# Prerequisites

* Thrift 0.9.3
```
$ thrift --version
Thrift version 0.9.3
```

* Golang and dep
```
$ dep version
dep:
 version     : v0.5.0
 build date  : 
 git hash    : 
 go version  : go1.10.2
 go compiler : gc
 platform    : darwin/amd64
 features    : ImportDuringSolve=false
```

# One Time Setup
## Initialize `dep`
```
loadtest_thrift093/go/src/simple_thrift_service$ GOPATH=`pwd`/../.. dep init
```

## Generate or Update Thrift files
```
loadtest_thrift093/go/src $ thrift --out . --gen go simple_thrift_service/thrift_def/hello.thrift
```

## Update the third party packages
```
loadtest_thrift093/go/src/simple_thrift_service$ GOPATH=`pwd`/../.. dep ensure
```

# Build and Run

## Compile
```
loadtest_thrift093/go/src/simple_thrift_service$ make
```

## Run Unit test
```
loadtest_thrift093/go/src/simple_thrift_service$ make test
```

## Run Server & Client
```
loadtest_thrift093/go$ ./bin/server --port 9394
2018/09/15 21:48:59 Starting server in localhost:9394
```
```
loadtest_thrift093/go$ ./bin/client --server_addr=localhost:9394 
2018/09/15 21:49:11 Talking to server at [localhost:9394]
2018/09/15 21:49:11 hello, world!
```

## Run Loadtest
After the server has been running.
```
loadtest_thrift093/go$ ./bin/loadtest --num_request=800 --qps=100
```
