#!/usr/bin/env bash
curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTex",
        "network": "mainnet"},
    "account_identifier": {
        "address": "io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"
    }}' http://127.0.0.1:8080/account/balance

curl -X POST --data '{"metadata": {}}' http://127.0.0.1:8080/network/list