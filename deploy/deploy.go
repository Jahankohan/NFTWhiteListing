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

	store "Go/nft_whitelisting/artifacts"
	u "Go/nft_whitelisting/utils"
)

func connect(dev string) *ethclient.Client {
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
	fmt.Println("Header Number:", header.Number.String())
	return client
}

func DeployContract(client *ethclient.Client) *store.Store{
	// This Method Deploy Contract 
	configuration := u.LoadConfig()
	pK := configuration.Local.PrivateKey

	privateKey, err := crypto.HexToECDSA(pK)
	if err!=nil {
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

	input := "1.0"
    address, tx, instance, err := store.DeployStore(auth, client, input)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Contract Deployed to: ", address.Hex())
    fmt.Println("TX Hash: ", tx.Hash().Hex())
	os.Setenv("CONTRACT_ADDRESS", address.Hex())
    return instance
}


func LoadContract() *store.Store{
	// This Method Handle Contract Deployment and Loading,
	// If Already Deployed, just return the instance
	// If not, it call DeployContract to Deploy it.
	client := connect("local")
	fmt.Println("Contract Address: ", os.Getenv("CONTRACT_ADDRESS"))
	if os.Getenv("CONTRACT_ADDRESS") == "" {
		fmt.Println("Not Deployed")
        instance := DeployContract(client)
		fmt.Println("Contract Address: ", os.Getenv("CONTRACT_ADDRESS"))
		return instance
	}
    address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
    instance, err := store.NewStore(address, client)
    if err != nil {
		log.Fatal(err)
    }

    fmt.Println("Contract is loaded")
    return instance
}

func QueryContract() {
    instance := LoadContract()
	fmt.Println(os.Getenv("CONTRACT_ADDRESS"))
    version, err := instance.Version(nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(version)
}

func WriteToContract() {
	client := connect("local")
	instance := LoadContract()
	configuration := u.LoadConfig()
	pK := configuration.Local.PrivateKey
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
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

    key := [32]byte{}
    value := [32]byte{}
    copy(key[:], []byte("foo"))
    copy(value[:], []byte("bar"))

    tx, err := instance.SetItem(auth, key, value)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

    result, err := instance.Items(nil, key)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(result[:]))
}
