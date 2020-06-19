#!/usr/bin/env bash
curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTex",
        "network": "mainnet",
        "sub_network_identifier": {}
    },
    "account_identifier": {
        "address": "io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg",
        "sub_account": {},
        "metadata": {}
    },
    "block_identifier": {}
}' http://127.0.0.1:8080/account/balance
