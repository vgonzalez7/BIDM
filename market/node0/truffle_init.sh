#!/usr/bin/env bash
geth --datadir ./ --syncmode full --nat=extip:127.0.0.1 --mine --miner.threads 1 --verbosity 3 --networkid 56424 --port 30300 --gasprice '0' --rpc --rpcport "8545" --rpccorsdomain "*" --bootnodes "enode://ff8cb1ac9c7eb398fedb08642033d2142b4ec248acb912d2e15f0201a000318b0f131a2c2c7ffe79385ecdeb7a9f5cbbce2e6c6e08cc2c9ffa9fc6324474c70e@127.0.0.1:0?discport=8009" --unlock 7fff1f93d22e9b2aa4c6a536531950ed386bd95e --password password.txt --allow-insecure-unlock
