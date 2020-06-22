#!/usr/bin/env bash
./rosetta-cli check --lookup-balance-by-block false --start 4032646 --end 4034781 --bootstrap-balances ./bootstrap_balances.json --block-concurrency 1

./rosetta-cli check --lookup-balance-by-block false --end 10 --bootstrap-balances ./bootstrap_balances.json --block-concurrency 1

./rosetta-cli view:account '{"address":"io10t7juxazfteqzjsd6qjk7tkgmngj2tm7n4fvrd"}'

./rosetta-cli view:block 4034780