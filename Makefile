.PHONY: proto build start

# Requires: protoc, protoc-gen-go, protoc-gen-go-grpc
# Install: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#          go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
proto:
	mkdir -p proto-files/generated-code
	find proto-files -name "*.proto" -exec protoc \
		--proto_path=proto-files \
		--go_out=proto-files/generated-code \
		--go_opt=paths=source_relative \
		--go-grpc_out=proto-files/generated-code \
		--go-grpc_opt=paths=source_relative \
		{} +

build:
	docker compose build

start: build
	docker compose up -d
