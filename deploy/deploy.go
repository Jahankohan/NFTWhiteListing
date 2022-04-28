package deploy

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	nft "Go/nft_whitelisting/artifacts"
	u "Go/nft_whitelisting/utils"
)

func Connect(dev string) *ethclient.Client {
	// Connecting to the Ganache
	if dev != "local" {
		log.Fatal("Other Methods Not Implemented Yet.")
	}
	configuration := u.LoadConfig()
	network := configuration.Local.Network
	client, err := ethclient.Dial(network)
	if err != nil {
		log.Fatal(err)
	}
	// Checking Connection
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	}
    _ = header
	return client
}

func GetAddedItemsCounter(instance *nft.NfthiteList, auth *bind.TransactOpts) {
    bind := bind.CallOpts{
        Pending: true,
        From:    auth.From,
        Context: auth.Context,
    }
    counter, err :=instance.AddedItems(&bind)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Current Counter", counter)
}

func DeployContract(client *ethclient.Client, max_allowed big.Int) (*nft.NfthiteList, error){
	// This Method Deploy Contract 
    auth := GetOwnerAuth()
    address, tx, instance, err := nft.DeployNfthiteList(auth, client, &max_allowed)
    if err != nil {
        log.Fatal(err)
        return instance, err
    }

    fmt.Println("Contract Deployed to: ", address.Hex())
    fmt.Println("TX Hash: ", tx.Hash().Hex())
	os.Setenv("CONTRACT_ADDRESS", address.Hex())
    return instance, err
}


func LoadContract(max_allowed big.Int) *nft.NfthiteList{
	// This Method Handle Contract Deployment and Loading,
	// If Already Deployed, just return the instance
	// If not, it call DeployContract to Deploy it.
	client := Connect("local")
	fmt.Println("Contract Address: ", os.Getenv("CONTRACT_ADDRESS"))
	if os.Getenv("CONTRACT_ADDRESS") == "" {
		fmt.Println("Not Deployed")
        instance, err := DeployContract(client, max_allowed)
        if err != nil{
            log.Fatal(err)
        }
		fmt.Println("Contract Address: ", os.Getenv("CONTRACT_ADDRESS"))
		return instance
	}
    address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
    instance, err := nft.NewNfthiteList(address, client)
    if err != nil {
		log.Fatal(err)
    }
    
    fmt.Println("Contract is loaded")
    return instance
}

func GetAuthOwnerKeys(pK string) *bind.TransactOpts{
    client := Connect("local")
    
    privateKey, err := crypto.HexToECDSA(pK)
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(3000000) // in units
    auth.GasPrice = gasPrice
    return auth;
}

func UpdateNonce(fromAddress common.Address) *big.Int{
    client := Connect("local")
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }
    return big.NewInt(int64(nonce));
}

func AddAddressToSC(instance *nft.NfthiteList, auth *bind.TransactOpts ,address []string) error{
	
    for _, item := range address{
        _ ,err := instance.AddAddressToWhiteList(auth, common.Hex2Bytes(item))
        auth.Nonce = UpdateNonce(auth.From)
        if err != nil{
            return err
        }
    }
    return nil
}

func GetListingRequest(instance *nft.NfthiteList, auth *bind.TransactOpts) {
    bind := bind.CallOpts{
        Pending: true,
        From:    auth.From,
        Context: auth.Context,
    }

    listing, err := instance.ListingRequests(&bind, big.NewInt(1))

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Listin Request 0:", listing)
}

func BuildMerkleTree(instance *nft.NfthiteList, auth *bind.TransactOpts) error{
    _, err := instance.BuildMerkleTree(auth)
    if err != nil{
        return err
    }
    return nil
}

func GetNotWhiteListedAuth() *bind.TransactOpts {
    configuration := u.LoadConfig()
	pK := configuration.Local.NonWhitelistedPrivateKey
    auth := GetAuthOwnerKeys(pK)
    return auth
}

func GetNotOwnerAuth() *bind.TransactOpts {
    configuration := u.LoadConfig()
	pK := configuration.Local.NonOwnerPrivateKey
    auth := GetAuthOwnerKeys(pK)
    return auth
}

func GetOwnerAuth() *bind.TransactOpts {
    configuration := u.LoadConfig()
	pK := configuration.Local.OwnerPrivateKey
    auth := GetAuthOwnerKeys(pK)
    return auth
}

func GetBindCallOps(auth *bind.TransactOpts) bind.CallOpts{

    bind := bind.CallOpts{
        Pending: true,
        From:    auth.From,
        Context: auth.Context,
    }
    return bind
}


func GetMerkleRoot(instance *nft.NfthiteList, bind *bind.CallOpts) [32]byte{
    root, err := instance.GetMerkleRoot(bind)

    if err != nil{
        log.Fatal(err)
    }
    
    return root
}

func IsWhiteListed(instance *nft.NfthiteList, auth *bind.TransactOpts, proof [][32]byte) bool {

    bind := bind.CallOpts{
        Pending: true,
        From:    auth.From,
        Context: auth.Context,
    }

    whiteListed, err := instance.SayHelloToWhiteListed(&bind, proof)
    if err != nil {
        return false
    }
    fmt.Println(whiteListed)
    return true
}

func GetProofForAddress(instance *nft.NfthiteList, auth *bind.TransactOpts) ([][32]byte, error){
    bind := bind.CallOpts{
        Pending: true,
        From:    auth.From,
        Context: auth.Context,
    }

    proof, err := instance.GetProof(&bind) 
    if err != nil{
        return proof, err
    }
    
    return proof, nil
}