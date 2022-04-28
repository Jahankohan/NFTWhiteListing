package main

import (
	deploy "Go/nft_whitelisting/deploy"
	"fmt"
	"log"
	"math/big"
)

func main() {
	// To test with this address just run: ganachi-cli -m "proud oxygen world expire ethics verb source fall reason donor grit lobster"
	// Or You can replace your own accounts, keep in mind to change config.yml file in /config

	max_allowed := big.NewInt(8)
	instance := deploy.LoadContract(*max_allowed)
	var address []string
	address = append(address, "f78DA8FAD6b18565C9553e01E3baF4136474f577")
	address = append(address, "0D6b48b0A49447C0ac125C6eA97B1D76956C8797")
	address = append(address, "39feC79D4e30f104e0367b9C378db8f474E4CE77")
	address = append(address, "a63A7369673c7a96aD9155b307826f2a3A0Dc7F1")
	address = append(address, "8c61fDF1facBa6B1972862E915B5F2166718a664")
	address = append(address, "Bd437ebDd2f3F061cE1892281a4fE4667DdFCEBC")
	address = append(address, "bF85850b16De96b04E3cDE1BB78D41baaCB941c3")
	address = append(address, "Fe7d26338ea24109C8540b98bE2cdc9145b11a44")
	auth := deploy.GetOwnerAuth()
	deploy.AddAddressToSC(instance, auth ,address)
	// Updating nonce after each execution
	auth.Nonce = deploy.UpdateNonce(auth.From)
	deploy.BuildMerkleTree(instance, auth)
	auth.Nonce = deploy.UpdateNonce(auth.From)
	proof, err := deploy.GetProofForAddress(instance, auth)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = deploy.UpdateNonce(auth.From)
	isWhiteListed := deploy.IsWhiteListed(instance, auth,proof)
	fmt.Println("User Is whiteListed!", isWhiteListed)
}

// ganachi-cli -m "proud oxygen world expire ethics verb source fall reason donor grit lobster"