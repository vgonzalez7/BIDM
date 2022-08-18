#!/usr/bin/env bash

geth --datadir ./ --syncmode full --verbosity 3 --networkid 56424 --port 30318 --nat=extip:10.10.150.230 --gasprice '0' --bootnodes "enode://b79a0254e23b49afb4fe99850770a5a826a95cbbb3a9a35f870c3d5684f78054021f1b47ea60864ad78e0aaf5c3624050d0d0effa224578fe67a0be2a1cee1b9@10.10.150.230:0?discport=8009"
