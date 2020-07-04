FROM golang:1.13.5-stretch
ENV GO111MODULE=on
WORKDIR apps/iotex-core-rosetta-gateway
COPY . .
CMD [ "cd tests","chmod +x test.sh","./test.sh"]