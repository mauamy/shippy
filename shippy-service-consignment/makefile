protoc:
	protoc -I. --go_out=plugins=grpc:. proto/consignment/consignment.proto

docker_build:
	docker build -t shippy-service-consignment .

docker_run: docker_build
	docker run -p 50051:50051 shippy-service-consignment