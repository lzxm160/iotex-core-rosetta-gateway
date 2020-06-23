# Running iotex-core-rosetta-gateway in Docker

This directory contains a Dockerfile that builds an image containing the [iotex-core-rosetta-gateway](https://github.com/iotexproject/iotex-core-rosetta-gateway).

To build the Docker image:

	docker build -f ./docker/Dockerfile . -t iotexproject/iotex-core-rosetta-gateway

To run the Docker image:

	docker run -p 8080:8080 -e "ConfigPath=/etc/iotex/config.yaml" iotexproject/iotex-core-rosetta-gateway

