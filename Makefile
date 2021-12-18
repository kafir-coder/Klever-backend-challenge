.PHONY: compile
compile: ## Compile the proto file.
	protoc -I pkg/proto/bank/account/ pkg/proto/bank/account/deposit.proto - 
.PHONY: build
build: ## build the go source code
	go build -race -ldflags "-s -w" -o bin/server cmd/server/main.go

.PHONY: run
run: ## Build and run server.
	bin/server 
