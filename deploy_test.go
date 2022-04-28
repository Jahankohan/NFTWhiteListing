package main

import (
	d "Go/nft_whitelisting/deploy"
	"math/big"
	"testing"
)

func TestDeployLocal(t *testing.T) {
	// Test Deployment of the Smart Contract
	client := d.Connect("local")
	_, err := d.DeployContract(client, *big.NewInt(8))
	if err != nil {
		t.Error("Contract Not Deployed! ", err)
	}
}

func TestAddAddressToWhiteListByOwner(t *testing.T){
	// Test Adding Address to white list by Owner
	client := d.Connect("local")
	var address []string
	// to test with this address just run: ganachi-cli -m "proud oxygen world expire ethics verb source fall reason donor grit lobster"
	address = append(address, "f78DA8FAD6b18565C9553e01E3baF4136474f577")
	address = append(address, "0D6b48b0A49447C0ac125C6eA97B1D76956C8797")
	address = append(address, "39feC79D4e30f104e0367b9C378db8f474E4CE77")
	address = append(address, "a63A7369673c7a96aD9155b307826f2a3A0Dc7F1")
	address = append(address, "8c61fDF1facBa6B1972862E915B5F2166718a664")
	address = append(address, "Bd437ebDd2f3F061cE1892281a4fE4667DdFCEBC")
	address = append(address, "bF85850b16De96b04E3cDE1BB78D41baaCB941c3")
	address = append(address, "Fe7d26338ea24109C8540b98bE2cdc9145b11a44")
	instance, err := d.DeployContract(client, *big.NewInt(8))
	if err != nil{
		t.Error("Failed to Initiat Contracts")
	}
	auth := d.GetOwnerAuth()

	err2 := d.AddAddressToSC(instance, auth, address)
	if err2 != nil{
		t.Error("Failed to add new address")
	}
}

func TestAddAddressToWhiteListByNonOwner(t *testing.T) {
	// Testing adding adddress to the white list with non owner account
	client := d.Connect("local")
	var address []string
	// to test with this address just run: ganachi-cli -m "proud oxygen world expire ethics verb source fall reason donor grit lobster"
	address = append(address, "f78DA8FAD6b18565C9553e01E3baF4136474f577")
	address = append(address, "0D6b48b0A49447C0ac125C6eA97B1D76956C8797")
	address = append(address, "39feC79D4e30f104e0367b9C378db8f474E4CE77")
	address = append(address, "a63A7369673c7a96aD9155b307826f2a3A0Dc7F1")
	address = append(address, "8c61fDF1facBa6B1972862E915B5F2166718a664")
	address = append(address, "Bd437ebDd2f3F061cE1892281a4fE4667DdFCEBC")
	address = append(address, "bF85850b16De96b04E3cDE1BB78D41baaCB941c3")
	address = append(address, "Fe7d26338ea24109C8540b98bE2cdc9145b11a44")
	instance, err := d.DeployContract(client, *big.NewInt(8))
	if err != nil{
		t.Error("Failed to Initiat Contracts")
	}

	auth := d.GetNotOwnerAuth()
	err2 := d.AddAddressToSC(instance, auth, address)
	if err2 == nil {
		t.Error("Only Admin Should be able to Add Address to smart contract.")
	}
}

func TestWhiteListed(t *testing.T) {
	// Testing to see if the provided proof works for listed account
	client := d.Connect("local")
	var address []string
	// to test with this address just run: ganachi-cli -m "proud oxygen world expire ethics verb source fall reason donor grit lobster"
	address = append(address, "f78DA8FAD6b18565C9553e01E3baF4136474f577")
	address = append(address, "0D6b48b0A49447C0ac125C6eA97B1D76956C8797")
	address = append(address, "39feC79D4e30f104e0367b9C378db8f474E4CE77")
	address = append(address, "a63A7369673c7a96aD9155b307826f2a3A0Dc7F1")
	address = append(address, "8c61fDF1facBa6B1972862E915B5F2166718a664")
	address = append(address, "Bd437ebDd2f3F061cE1892281a4fE4667DdFCEBC")
	address = append(address, "bF85850b16De96b04E3cDE1BB78D41baaCB941c3")
	address = append(address, "Fe7d26338ea24109C8540b98bE2cdc9145b11a44")

	instance, err := d.DeployContract(client, *big.NewInt(8))
	if err != nil{
		t.Error("Failed to Initiat Contracts")
	}

	auth := d.GetOwnerAuth()

	err2 := d.AddAddressToSC(instance, auth, address)
	if err2 != nil{
		t.Error("Failed to add new address")
	}
	auth.Nonce = d.UpdateNonce(auth.From)
	err3 := d.BuildMerkleTree(instance, auth)
	if err3 != nil{
		t.Error("Failed to build merkle tree")
	}
	auth.Nonce = d.UpdateNonce(auth.From)
	proof, err := d.GetProofForAddress(instance, auth)
	if err != nil {
		t.Error("Faild to Calculate the Proof")
	}
	auth.Nonce = d.UpdateNonce(auth.From)
	isWhiteListed := d.IsWhiteListed(instance, auth, proof)

	if isWhiteListed != true {
		t.Error("Faild to detect white listed user")
	}

}

func TestNonWhiteListed(t *testing.T) {
	// test to see if white list reject the non whitelisted account
	client := d.Connect("local")
	var address []string
	// to test with this address just run: ganachi-cli -m "proud oxygen world expire ethics verb source fall reason donor grit lobster"
	address = append(address, "f78DA8FAD6b18565C9553e01E3baF4136474f577")
	address = append(address, "0D6b48b0A49447C0ac125C6eA97B1D76956C8797")
	address = append(address, "39feC79D4e30f104e0367b9C378db8f474E4CE77")
	address = append(address, "a63A7369673c7a96aD9155b307826f2a3A0Dc7F1")
	address = append(address, "8c61fDF1facBa6B1972862E915B5F2166718a664")

	instance, err := d.DeployContract(client, *big.NewInt(8))
	if err != nil{
		t.Error("Failed to Initiat Contracts")
	}

	auth := d.GetOwnerAuth()

	err2 := d.AddAddressToSC(instance, auth, address)
	if err2 != nil{
		t.Error("Failed to add new address")
	}
	auth.Nonce = d.UpdateNonce(auth.From)
	err3 := d.BuildMerkleTree(instance, auth)
	if err3 != nil{
		t.Error("Failed to build merkle tree")
	}
	auth.Nonce = d.UpdateNonce(auth.From)
	proof, err := d.GetProofForAddress(instance, auth)
	if err != nil {
		t.Error("Faild to Calculate the Proof")
	}
	nonListedAuth := d.GetNotWhiteListedAuth()
	isWhiteListed := d.IsWhiteListed(instance, nonListedAuth, proof)

	if isWhiteListed == true {
		t.Error("Faild to detect non white listed user")
	}

}