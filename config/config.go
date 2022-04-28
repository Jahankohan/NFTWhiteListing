package config

type Configurations struct {
	Local       LocalConfigurations
	Testnet     TestnetConfigurations
	Mainnet 	MainnetConfigurations
}

type LocalConfigurations struct {
	OwnerPrivateKey	string
	NonOwnerPrivateKey	string
	NonWhitelistedPrivateKey	string
	Network		string
	ChainId		int
}


type TestnetConfigurations struct {
	OwnerPrivateKey	string
	NonOwnerPrivateKey	string
	NonWhitelistedPrivateKey	string
	PrivateKey	string
	Network		string
}


type MainnetConfigurations struct {
	OwnerPrivateKey	string
	NonOwnerPrivateKey	string
	NonWhitelistedPrivateKey	string
	PrivateKey	string
	Network		string
}