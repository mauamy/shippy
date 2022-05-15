protoc:
	 protoc --proto_path=. --go_out=. --micro_out=. proto/vessel/vessel.proto

docker_build: protoc
	docker build -t shippy-service-vessel .

docker_run:
	docker run shippy-service-vessel