GOCMD=go
GOPATH=$(dir $(abspath ..))
OUTPUT_DIR=$(GOPATH)bin

build:
	@echo "GOPATH=$(GOPATH)"
	cd $(GOPATH)
	mkdir -p $(OUTPUT_DIR)

	@GOPATH=$(GOPATH) $(GOCMD) build -v -o $(OUTPUT_DIR)/metrics simple_thrift_service/cmd/metrics/$(wildcard *.go)
	@echo $(OUTPUT_DIR)/metrics

	@GOPATH=$(GOPATH) $(GOCMD) build -v -o $(OUTPUT_DIR)/server simple_thrift_service/cmd/server/$(wildcard *.go)
	@echo $(OUTPUT_DIR)/server

	@GOPATH=$(GOPATH) $(GOCMD) build -v -o $(OUTPUT_DIR)/client simple_thrift_service/cmd/client/$(wildcard *.go)
	@echo $(OUTPUT_DIR)/client

	@GOPATH=$(GOPATH) $(GOCMD) build -v -o $(OUTPUT_DIR)/loadtest simple_thrift_service/cmd/loadtest/$(wildcard *.go)
	@echo $(OUTPUT_DIR)/loadtest

test:
	cd $(GOPATH)
	@GOPATH=$(GOPATH) $(GOCMD) test -v simple_thrift_service/service/...

clean:
	rm -rf $(OUTPUT_DIR)
