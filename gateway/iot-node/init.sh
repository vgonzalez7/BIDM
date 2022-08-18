#!/usr/bin/env bash

geth --datadir ./ --syncmode full --verbosity 3 --networkid 56424 --port 30319 --nat=extip:10.10.150.229 --gasprice '0'  --bootnodes "enode://ff8cb1ac9c7eb398fedb08642033d2142b4ec248acb912d2e15f0201a000318b0f131a2c2c7ffe79385ecdeb7a9f5cbbce2e6c6e08cc2c9ffa9fc6324474c70e@10.10.150.230:0?discport=8009"
