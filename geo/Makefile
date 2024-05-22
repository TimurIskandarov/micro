OS := $(shell uname -o)
LOCAL_BIN := "$(CURDIR)/bin"

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=54322 dbname=geo user=geo-user password=geo-password sslmode=disable"

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
	mkdir -p pkg/geo_v1
	protoc --proto_path api/geo_v1 \
	--go_out=pkg/geo_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(GO_PLUGIN) \
	--go-grpc_out=pkg/geo_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(GRPC_PLUGIN) \
	api/geo_v1/service.proto

local-migration-status:
	goose -dir ${LOCAL_NIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v
	
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v