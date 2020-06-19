#!/usr/bin/env bash
curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTex",
        "network": "mainnet"},
    "account_identifier": {
        "address": "io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"
    }}' http://127.0.0.1:8080/account/balance

#response:
#{"block_identifier":{"index":3986321,"hash":"931345a809f68dd454716f75c3a08350232be071a56212fac7fb666fc4e608c5"},"balances":[{"value":"12000000000000000000","currency":{"symbol":"IoTex","decimals":18}}],"metadata":{"nonce":0}}


curl -X POST --data '{"metadata": {}}' http://127.0.0.1:8080/network/list
#response:
#{"network_identifiers":[{"blockchain":"IoTex","network":"mainnet"}]}

curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTex",
        "network": "mainnet"},
    "block_identifier": {"index": 390873}}' http://127.0.0.1:8080/block
#response:
#{"block":{"block_identifier":{"index":390873,"hash":"5c084459315fcf0839ed9f2d8b89ca8fb039695a56007a071e5ce9d3c8908d95"},"parent_block_identifier":{"index":390872,"hash":"3ae76de97535f4908d7dd6b2d5f232543b1e5a9fe80a0e9d8f91fdd27d9363eb"},"timestamp":1573620900000,"transactions":[{"transaction_identifier":{"hash":"b37d5db44bd3dc182617b56744e12cab94486808eae1dc401599b611ed388164"},"operations":[{"operation_identifier":{"index":9},"type":"transfer","status":"succeed","account":{"address":"io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"},"amount":{"value":"10000000000000000000","currency":{"symbol":"IoTex","decimals":18}}}]}]}}

curl -X POST --data '{
    "network_identifier": {
        "blockchain": "bitcoin",
        "network": "mainnet",
        "sub_network_identifier": {
            "network": "shard 1",
            "metadata": {
                "producer": "0x52bc44d5378309ee2abf1539bf71de1b7d7be3b5"
            }
        }
    },
    "options": {}
}' http://127.0.0.1:8080/construction/metadata

