#!/usr/bin/env bash
go build -o ./server .
export IoTexChainPoint=api.testnet.iotex.one:80
./server
