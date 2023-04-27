protoc --go_out=. --go_opt=paths=source_relative \
 --go-grpc_out=. --go-grpc_opt=paths=source_relative \
 proto/\*.proto

protoc --plugin=protoc-gen-grpc-java --java_out=./java --java-grpc_out=./java proto/greeter.proto
