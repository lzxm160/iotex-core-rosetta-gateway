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

GW="./server -config-path=config_testnet.yaml -genesis-path=genesis_testnet.yaml -plugin=gateway"
printf "${GRN}### Starting the iotex server...${OFF}\n"
${GW} &

sleep 3

printf "${GRN}### Inject some actions...${OFF}\n"
cd inject
go test

sleep 3 #wait for the last candidate action

cd ..
GW="./iotex-core-rosetta-gateway"
printf "${GRN}### Starting the Rosetta gateway...${OFF}\n"
cd ..
export ConfigPath=tests/gateway_config.yaml
go build -o ./iotex-core-rosetta-gateway .
${GW} &

sleep 3

printf "${GRN}### Run rosetta-cli create:configuration...${OFF}\n"
cd tests
./rosetta-cli create:configuration config.json

printf "${GRN}### Validating Rosetta gateway implementation...${OFF}\n"
./rosetta-cli check --lookup-balance-by-block=false --end 111 --bootstrap-balances ./bootstrap_balances.json --block-concurrency 8
./rosetta-cli view:account '{"address":"io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"}'
./rosetta-cli view:block 10
rm -rf /tmp/rosetta-cli*

# Clean up after a successful run.
rm -rf ./rosetta* ./iotex-core* ./*.db ./server ./*.tar.gz
rm -rf ./config.json
printf "${GRN}### Tests finished.${OFF}\n"