#!/usr/bin/env bash
1. curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTex",
        "network": "mainnet"},
    "account_identifier": {
        "address": "io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"
    }}' http://127.0.0.1:8080/account/balance

response:
{"block_identifier":{"index":3986321,"hash":"931345a809f68dd454716f75c3a08350232be071a56212fac7fb666fc4e608c5"},"balances":[{"value":"12000000000000000000","currency":{"symbol":"IoTex","decimals":18}}],"metadata":{"nonce":0}}


2. curl -X POST --data '{"metadata": {}}' http://127.0.0.1:8080/network/list
response:
{"network_identifiers":[{"blockchain":"IoTex","network":"mainnet"}]}

3. curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTex",
        "network": "mainnet"},
    "block_identifier": {"index": 1123941}}' http://127.0.0.1:8080/block