OS := $(shell uname -o)
LOCAL_BIN := "$(CURDIR)/bin"

ifeq ($(OS), Msys)
	GO_PLUGIN=bin/protoc-gen-go.exe
	GRPC_PLUGIN=bin/protoc-gen-go-grpc.exe
else
	GO_PLUGIN=bin/protoc-gen-go
	GRPC_PLUGIN=bin/protoc-gen-go-grpc
endif

install-go-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

generate:
	mkdir -p pkg/auth_v1
	protoc --proto_path api/auth_v1 \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(GO_PLUGIN) \
	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(GRPC_PLUGIN) \
	api/auth_v1/service.proto
