
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
```
loadtest_thrift093/go/src/simple_thrift_service$ GOPATH=`pwd`/../.. dep init
```

```
loadtest_thrift093/go/src/simple_thrift_service$ GOPATH=`pwd`/../.. dep ensure
```
