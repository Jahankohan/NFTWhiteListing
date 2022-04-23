package config

type Configurations struct {
	Local       LocalConfigurations
	Testnet     TestnetConfigurations
	Mainnet 	MainnetConfigurations
}

type LocalConfigurations struct {
	PrivateKey	string
	Network		string
	ChainId		int
}


type TestnetConfigurations struct {
	PrivateKey	string
	Network		string
}


type MainnetConfigurations struct {
	PrivateKey	string
	Network		string
}