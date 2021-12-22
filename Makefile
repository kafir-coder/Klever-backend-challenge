.PHONY: compile
compile: ## Compile the proto file.
	protoc -I pkg/proto/token -I ${GOPATH}src -I ${GOPATH}src/github.com/envoyproxy/protoc-gen-validate pkg/proto/token/token.proto --go_out=plugins=grpc:pkg/proto/token $(find /pkg/proto/token -name '*.proto')

.PHONY: build
build: ## build the go source code
	go build -race -ldflags "-s -w" -o bin/server cmd/server/main.go

.PHONY: run
run: ## Build and run server.
	bin/server 

.PHONY: test
test: ## run the test cases
	go test -v -run ./ ./internal/token
