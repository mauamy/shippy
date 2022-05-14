build_service_consignment_docker:
	cd shippy-service-consignment && $(MAKE) docker_build

run_service_consignment_docker:
	cd shippy-service-consignment && $(MAKE) docker_run

### ------------------ ###

build_cli_consignment_docker:
	docker build -f shippy-cli-consignment/Dockerfile -t shippy-cli-consignment .

run_cli_consignment_docker:
	cd shippy-cli-consignment && $(MAKE) docker_run
