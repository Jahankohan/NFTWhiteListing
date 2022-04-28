docker run --rm -v $(pwd):/root ethereum/solc:0.8.1 --bin /root/contracts/NFTWhiteList.sol -o /root/artifacts/bin
docker run --rm -v $(pwd):/root ethereum/solc:0.8.1 --abi /root/contracts/NFTWhiteList.sol -o /root/artifacts/build
abigen --bin=artifacts/bin/NFTWhiteList.bin --abi=artifacts/build/NFTWhiteList.abi --pkg=nfthiteList --out=artifacts/WhiteList.go