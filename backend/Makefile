.PHONY: proto

proto:
	protoc --go_out=pb \
		--go-grpc_out=pb \
		--go-grpc_opt require_unimplemented_servers=false \
		proto/sample.proto

