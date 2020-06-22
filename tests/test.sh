#!/usr/bin/env bash
# ANSI escape codes to brighten up the output.
GRN=$'\e[32;1m'
OFF=$'\e[0m'
GW="./iotex-core-rosetta-gateway"
printf "${GRN}### Starting the Rosetta gateway...${OFF}\n"
${GW} &

sleep 3

printf "${GRN}### Validating Rosetta gateway implementation...${OFF}\n"
./rosetta-cli check --lookup-balance-by-block false --end 10 --bootstrap-balances ./bootstrap_balances.json --block-concurrency 1
./rosetta-cli view:account '{"address":"io10t7juxazfteqzjsd6qjk7tkgmngj2tm7n4fvrd"}'
./rosetta-cli view:block 4034780
rm -rf /tmp/rosetta-cli*

# Clean up after a successful run.
rm -rf ./test/rosetta*

printf "${GRN}### Tests finished.${OFF}\n"