proto_go:
	protoc -I ./proto/ ./proto/*.proto --go_out=admin-service/gen --go-grpc_out=admin-service/gen

proto_py:
   python3 -m grpc_tools.protoc -I./proto --python_out=./cluster-service/gen \
   --pyi_out=./cluster-service/gen --grpc_python_out=./cluster-service/gen ./proto/app_auth.proto \ 
