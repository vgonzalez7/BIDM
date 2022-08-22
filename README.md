## BIDM

This repository contains the code to replicate the Blockchain-based IoT Data Marketplace. The steps necessary to get everything running are as follows:

- Make sure you have all dependencies - IPFS, NodeJS, and Go-Ethereum
- Start the bootnode (you can use the shell script in the market folder)
- Start the other nodes -node0 in market, clients-node in market and iot-node in gateway (it's also possible to use shell scripts within each node's folder)
- Run the market core (executable in market/market-module/core)
- Run the IoT proxy (executable in gateway/iot-proxy)
- Run the web API (node app in market/nodejs-client-server)

You can now interact with the Smart Contracts through CLI or tools like Remix, or use the web API to have a look at the transactions. By default the web API runs on port 5054.
