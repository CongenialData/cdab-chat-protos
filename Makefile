
# Define the input and output directories
PROTO_DIR := protos
OUTPUT_DIR := pkg/pb

# Define the protobuf compiler command
PROTOC := protoc

# Define the protobuf compiler flags
PROTOC_FLAGS := --go_out=$(OUTPUT_DIR) --go_opt=paths=source_relative \
	--go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=paths=source_relative \
	--proto_path=$(PROTO_DIR)

# Define the list of protobuf files to compile
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Define the target for compiling protobuf files
.PHONY: protos
protos:
	$(PROTOC) $(PROTOC_FLAGS) $(PROTO_FILES)



# .PHONY: protos
# 
# protos: protos/cdab-chat.proto
# 	protoc --go_out=pkg/protos --go_opt=paths=source_relative \
#     --go-grpc_out=pkg/protos --go-grpc_opt=paths=source_relative \
#     --proto_path=protos \
#     protos/*.proto%   