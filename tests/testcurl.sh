#!/usr/bin/env bash
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
    "account_identifier": {
        "address": "0x3a065000ab4183c6bf581dc1e55a605455fc6d61",
        "sub_account": {
            "address": "0x6b175474e89094c44da98b954eedeac495271d0f",
            "metadata": {}
        },
        "metadata": {}
    },
    "block_identifier": {
        "index": 1123941,
        "hash": "0x1f2cc6c5027d2f201a5453ad1119574d2aed23a392654742ac3c78783c071f85"
    }
}' http://127.0.0.1:8080/account/balance
