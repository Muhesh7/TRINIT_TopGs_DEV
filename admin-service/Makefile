build:
	go run /app/main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

proto:
	protoc -I ./protos/ ./protos/*.proto --go_out=./gen --go-grpc_out=./gen