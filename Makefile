.PHONY: compile
compile: ## Compile the proto file.
	protoc -I pkg/proto/token pkg/proto/token/token.proto --go_out=plugins=grpc:pkg/proto/token
.PHONY: build
build: ## build the go source code
	go build -race -ldflags "-s -w" -o bin/server cmd/server/main.go

.PHONY: run
run: ## Build and run server.
	bin/server 
