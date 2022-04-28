# NFTWhiteListing

White Listing mechanism implemented Using Solidity and Go, and for this purpose we used Solidity to implement Merkle Tree, Go is used for Deploying, Interacting and Testing Smart Contract.

## Compiling
The Artifacts directory contains compiled smart contract Using Abigen and Solc (You need docker to be run for this), the ABI, BIN and also Go interface. In case of any changes to the smart contract you should compile it again, there is a compile.sh file in root directory, just run in to get new compiled data! If artifacts directory is not empty, use `--overwrite` flag to force it or just simply delete files. 

## Network
For test purposes I used Ganache, If you want the same account to work just use bellow command

`ganachi-cli -m "proud oxygen world expire ethics verb source fall reason donor grit lobster"`

or you can run in another enviroments and just change address, and make sure to change privatekyes in /config/config.yml too.

## Test
For testing smart contract, just type the bellow code in root directory of the project:

`go test -v`
