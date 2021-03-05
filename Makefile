
BUF_CACHE_DIR=$(HOME)/.cache/buf
GOOGLEAPI_DIR=$(BUF_CACHE_DIR)/mod/buf.build/beta/googleapis/1dc4674e3cb949b388204fa2dc321be7
PHP_OUT=sdk/src

.PHONY: generate-php
generate-php:
	mkdir -p $(PHP_OUT)
	find $(GOOGLEAPI_DIR)/google/api -name "*.proto" -exec \
		protoc -I $(GOOGLEAPI_DIR) --php_out=./$(PHP_OUT) --grpc_out=./$(PHP_OUT) \
    --plugin=protoc-gen-grpc=./grpc_php_plugin {} \;
	find ./proto -name "*.proto" -exec \
		protoc -I ./proto -I $(GOOGLEAPI_DIR) --php_out=./$(PHP_OUT) --grpc_out=./$(PHP_OUT) \
    --plugin=protoc-gen-grpc=./grpc_php_plugin {} \;

.PHONY: generate-go
generate-go:
	buf generate

generate: generate-go generate-php
	@echo
	@echo generate OK
	@echo

lint:
	buf lint
	buf breaking --against 'https://github.com/johanbrandhorst/grpc-gateway-boilerplate.git#branch=master'

install: install-go install-buf

.PHONY: install-buf
install-buf:
	mkdir -p $(HOME)/.cache
	tar -zxf .data/buf.generate.tar.gz -C $(HOME)/.cache


.PHONY: install-go
install-go:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go@latest \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
		github.com/bufbuild/buf/cmd/buf@latest

PHP_DEPENDENCES=autoconf automake libtool shtool bazel
.PHONY: install-php
install-php:
	$(foreach exec,$(PHP_DEPENDENCES),$(if $(shell which $(exec)),,$(error "No bazel in $(PATH), consider doing install $(exec)")))
	git clone --recursive --depth 1 --shallow-submodules -b v1.35.0 https://github.com/grpc/grpc
	cd grpc
	#bazel build :all
	bazel build src/compiler:grpc_php_plugin
	cd ..
	cp grpc/bazel-bin/src/compiler/grpc_php_plugin ./
	rm -rf grpc
