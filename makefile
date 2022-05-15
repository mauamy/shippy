build_docker_all: build_cli_consignment_docker build_service_consignment_docker build_service_vessel_docker

build_service_consignment_docker:
	docker build -f shippy-service-consignment/Dockerfile -t shippy-service-consignment .

run_service_consignment_docker:
	cd shippy-service-consignment && $(MAKE) docker_run

### ------------------ ###

build_cli_consignment_docker:
	docker build -f shippy-cli-consignment/Dockerfile -t shippy-cli-consignment .

run_cli_consignment_docker:
	cd shippy-cli-consignment && $(MAKE) docker_run


### ------------------ ###
build_service_vessel_docker:
	cd shippy-service-vessel && $(MAKE) docker_build

run_service_vessel_docker:
	cd shippy-service-vessel && $(MAKE) docker_run
