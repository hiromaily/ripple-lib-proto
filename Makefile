PWD=$(shell pwd)
UID=$(shell id -u)
UNAME=$(shell whoami)
GNAME=$(shell id -gn)

PROTOC_BIN=protoc

GOGO_MODULES=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types

###############################################################################
# Setup
###############################################################################
.PHONY:install
install:
	brew install protobuf
	brew install prototool

.PHONY:install-plugin
install-plugin:
	# refer to https://developers.google.com/protocol-buffers/docs/reference/go-generated
	# provides a protoc-gen-go binary which protoc uses when invoked with the --go_out command-line flag.
	# The --go_out flag
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	# The go-grpc_out flag
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	# The --gofast_out flag
	go install github.com/gogo/protobuf/protoc-gen-gofast@latest
	# The --gogoslick_out flag
	go install github.com/gogo/protobuf/protoc-gen-gogoslick@latest

.PHONY:proto-lint
proto-lint:
	prototool lint

.PHONY:get-proto
get-proto:
	./scripts/get_proto.sh

# protoc command
.PHONY:protoc-go
protoc-go: clean
	$(PROTOC_BIN) \
	--go_out=./pb/go/rippleapi/ --go_opt=paths=source_relative \
	--go-grpc_out=./pb/go/rippleapi/ --go-grpc_opt=paths=source_relative  \
	--proto_path=./proto/rippleapi \
	--proto_path=./proto/third_party \
	proto/**/*.proto

#	$(PROTOC_BIN) --proto_path=proto -I=. \
#	--gogoslick_out=$(GOGO_MODULES):./pb/go/ proto/**/*.proto
#	chown -R $(UNAME):$(GNAME) ./pb/go

clean:
	rm -rf pb/go/rippleapi/*.go
