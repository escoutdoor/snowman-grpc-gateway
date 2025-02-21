LOCAL_BIN := $(CURDIR)/bin

VENDOR_PROTO_PATH := $(CURDIR)/vendor.protobuf

generate: .install-deps .buf-generate

.install-deps: export GOBIN := $(LOCAL_BIN)
.install-deps:
	go install github.com/bufbuild/buf/cmd/buf@v1.41.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.22.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.22.0

.buf-generate:
	./bin/buf generate

vendor:	.vendor-reset .vendor-googleapis .vendor-google-protobuf .vendor-protovalidate .vendor-protoc-gen-openapiv2 .vendor-tidy

# delete VENDOR_PROTO_PATH
.vendor-reset:
	rm -rf $(VENDOR_PROTO_PATH)
	mkdir -p $(VENDOR_PROTO_PATH)

# google/protobuf
.vendor-google-protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf $(VENDOR_PROTO_PATH)/protobuf &&\
	cd $(VENDOR_PROTO_PATH)/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p $(VENDOR_PROTO_PATH)/google/protobuf
	find $(VENDOR_PROTO_PATH)/protobuf/src/google/protobuf -maxdepth 1 -type f -exec mv {} $(VENDOR_PROTO_PATH)/google/protobuf \;
	rm -rf $(VENDOR_PROTO_PATH)/protobuf

# google/api
.vendor-googleapis:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/googleapis/googleapis $(VENDOR_PROTO_PATH)/googleapis &&\
	cd $(VENDOR_PROTO_PATH)/googleapis &&\
	git sparse-checkout set --no-cone google/api &&\
	git checkout
	mkdir -p $(VENDOR_PROTO_PATH)/google/api
	find $(VENDOR_PROTO_PATH)/googleapis/google/api -maxdepth 1 -type f -exec mv {} $(VENDOR_PROTO_PATH)/google/api \;
	rm -rf $(VENDOR_PROTO_PATH)/googleapis

# protoc-gen-openapiv2/options
.vendor-protoc-gen-openapiv2:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway $(VENDOR_PROTO_PATH)/grpc-gateway && \
 	cd $(VENDOR_PROTO_PATH)/grpc-gateway && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p $(VENDOR_PROTO_PATH)/protoc-gen-openapiv2
	mv $(VENDOR_PROTO_PATH)/grpc-gateway/protoc-gen-openapiv2/options $(VENDOR_PROTO_PATH)/protoc-gen-openapiv2
	rm -rf $(VENDOR_PROTO_PATH)/grpc-gateway

# buf/validate/validate.proto
.vendor-protovalidate:
	git clone -b main --single-branch --depth=1 --filter=tree:0 \
		https://github.com/bufbuild/protovalidate $(VENDOR_PROTO_PATH)/protovalidate && \
	cd $(VENDOR_PROTO_PATH)/protovalidate
	mv $(VENDOR_PROTO_PATH)/protovalidate/proto/protovalidate/buf $(VENDOR_PROTO_PATH)
	rm -rf $(VENDOR_PROTO_PATH)/protovalidate

# delete all non .proto files
.vendor-tidy:
	find $(VENDOR_PROTO_PATH) -type f ! -name "*.proto" -delete
	find $(VENDOR_PROTO_PATH) -type f \( -name "*unittest*.proto" -o -name "*test*.proto" -o -name "*sample*.proto" \) -delete
	find $(VENDOR_PROTO_PATH) -empty -type d -delete

run:
	@go run ./cmd/snowman/main.go
