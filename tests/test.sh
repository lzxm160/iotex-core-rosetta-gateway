#!/usr/bin/env bash
set -o nounset -o pipefail -o errexit

# Kill all dangling processes on exit.
cleanup() {
	printf "${OFF}"
	pkill -P $$ || true
}
trap "cleanup" EXIT

# ANSI escape codes to brighten up the output.
GRN=$'\e[32;1m'
OFF=$'\e[0m'

ROSETTA_CLI_RELEASE=0.4.1
# TODO change it to some version
IOTEX_SERVER_RELEASE=v1.1.0

export GO111MODULE=on
git config --global http.proxy 'socks5://192.168.1.8:1080'
git config --global https.proxy 'socks5://192.168.1.8:1080'
export http_proxy=socks5://192.168.1.8:1080

#download
printf "${GRN}### Downloading rosetta-cli release %s...${OFF}\n" ${ROSETTA_CLI_RELEASE}
	wget --quiet --show-progress --progress=bar:force:noscroll -O tests/rosetta-cli-${ROSETTA_CLI_RELEASE}.tar.gz https://github.com/coinbase/rosetta-cli/archive/v${ROSETTA_CLI_RELEASE}.tar.gz

printf "${GRN}### Downloading iotex-core release ${IOTEX_SERVER_RELEASE}...${OFF}\n"
	wget --quiet --show-progress --progress=bar:force:noscroll -O tests/iotex-core-${IOTEX_SERVER_RELEASE}.tar.gz https://github.com/iotexproject/iotex-core/archive/${IOTEX_SERVER_RELEASE}.tar.gz

#build
printf "${GRN}### Building rosetta-cli...${OFF}\n"
	tar -xf tests/rosetta-cli-${ROSETTA_CLI_RELEASE}.tar.gz -C tests
	cd tests/rosetta-cli-${ROSETTA_CLI_RELEASE} && go build
	cd ../..
	cp tests/rosetta-cli-${ROSETTA_CLI_RELEASE}/rosetta-cli tests

printf "${GRN}### Building iotex-core...${OFF}\n"
	tar -xf tests/iotex-core-${IOTEX_SERVER_RELEASE}.tar.gz -C tests
	cd tests/iotex-core-${IOTEX_SERVER_RELEASE} && make build
	cd ../..
	cp tests/iotex-core-${IOTEX_SERVER_RELEASE}/bin/server tests

cd tests
printf "${GRN}### Starting the iotex server...${OFF}\n"
GW="iotex-server -config-path=config_testnet.yaml -genesis-path=genesis_testnet.yaml -plugin=gateway"
${GW} &
sleep 3

printf "${GRN}### Starting the Rosetta gateway...${OFF}\n"
GW="iotex-core-rosetta-gateway"
${GW} &
sleep 3

cd ../rosetta-cli-config
printf "${GRN}### Run rosetta-cli check...${OFF}\n"
rosetta-cli check:data --configuration-file testing/iotex-testing.json &

cd ../tests/inject
printf "${GRN}### Inject some actions...${OFF}\n"
go test

sleep 10 #wait for the last candidate action

cd ../../rosetta-cli-config
printf "${GRN}### Run rosetta-cli view:account and view:block...${OFF}\n"
rosetta-cli view:account '{"address":"io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"}' --configuration-file testing/iotex-testing.json
rosetta-cli view:block 10 --configuration-file testing/iotex-testing.json

printf "${GRN}### Tests finished.${OFF}\n"
