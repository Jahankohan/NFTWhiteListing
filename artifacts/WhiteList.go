// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nfthiteList

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NfthiteListMetaData contains all meta data concerning the NfthiteList contract.
var NfthiteListMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max_allowed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addr\",\"type\":\"bytes\"}],\"name\":\"AddAddressToWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BuildMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"SayHelloToWhiteListed\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addedItems\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listingRequests\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260006003553480156200001657600080fd5b50604051620022423803806200224283398181016040528101906200003c9190620001bd565b33600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006002826200008e919062000241565b14620000a5578080620000a190620001f3565b9150505b8067ffffffffffffffff811115620000e6577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015620001155781602001602082028036833780820191505090505b50600090805190602001906200012d92919062000135565b5050620002f1565b82805482825590600052602060002090810192821562000174579160200282015b828111156200017357825182559160200191906001019062000156565b5b50905062000183919062000187565b5090565b5b80821115620001a257600081600090555060010162000188565b5090565b600081519050620001b781620002d7565b92915050565b600060208284031215620001d057600080fd5b6000620001e084828501620001a6565b91505092915050565b6000819050919050565b60006200020082620001e9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141562000236576200023562000279565b5b600182019050919050565b60006200024e82620001e9565b91506200025b83620001e9565b9250826200026e576200026d620002a8565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b620002e281620001e9565b8114620002ee57600080fd5b50565b611f4180620003016000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806387f8151c1161006657806387f8151c146100fc5780638da5cb5b1461012c578063a04e63bb1461014a578063a4730acf1461017a578063e3ad6acf1461019857610093565b80632eb4a7ab146100985780635c5d625e146100b65780637415ad99146100d45780637d3e9a78146100f2575b600080fd5b6100a06101b4565b6040516100ad919061195b565b60405180910390f35b6100be6101ba565b6040516100cb9190611939565b60405180910390f35b6100dc61024f565b6040516100e99190611a18565b60405180910390f35b6100fa610255565b005b61011660048036038101906101119190611624565b61084f565b6040516101239190611976565b60405180910390f35b610134610901565b604051610141919061191e565b60405180910390f35b610164600480360381019061015f91906116a6565b610927565b604051610171919061195b565b60405180910390f35b61018261094b565b60405161018f919061195b565b60405180910390f35b6101b260048036038101906101ad9190611665565b6109e5565b005b60015481565b60606000336040516020016101cf91906118a5565b6040516020818303038152906040528051906020012090506004600082815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561024457602002820191906000526020600020905b815481526020019060010190808311610230575b505050505091505090565b60035481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102dc906119f8565b60405180910390fd5b6000806000808060001b9050600060026003546103029190611d0b565b1461036f576003600081548092919061031a90611c94565b919050555060008160405160200161033291906118c0565b6040516020818303038152906040528051906020012090806001815401808255809150506001900390600052602060002001600090919091909150555b61037a600354610b4a565b809450819550505060008467ffffffffffffffff8111156103c4577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156103f25781602001602082028036833780820191505090505b50905060005b6003548110156104ad576000818154811061043c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154828281518110610480577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018181525050838061049790611c94565b94505080806104a590611c94565b9150506103f8565b5060008060008590505b878610156107e6578192505b8083101561073e576000846001856104db9190611b20565b81518110610512577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151858581518110610553577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151101561061d5784848151811061059a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151856001866105b09190611b20565b815181106105e7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040516020016106009291906118db565b6040516020818303038152906040528051906020012090506106d5565b8460018561062b9190611b20565b81518110610662577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518585815181106106a3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040516020016106bc9291906118db565b6040516020818303038152906040528051906020012090505b8085888151811061070f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018181525050868061072690611c94565b9750506002846107369190611b20565b9350506104c3565b80915060006002876107509190611d0b565b1415801561075e5750878614155b156107de578460405160200161077491906118c0565b604051602081830303815290604052805190602001208487815181106107c3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101818152505085806107da90611c94565b9650505b8590506104b7565b61083b8460018a6107f79190611ba7565b8151811061082e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151610bed565b6108458488610bf7565b5050505050505050565b60606000801b6001541415610899576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610890906119b8565b60405180910390fd5b6108a282611288565b6108e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d8906119d8565b60405180910390fd5b6040518060600160405280602c8152602001611ee0602c91399050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000818154811061093757600080fd5b906000526020600020016000915090505481565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d4906119f8565b60405180910390fd5b600154905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6c906119f8565b60405180910390fd5b60008054905060035410610abe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ab590611998565b60405180910390fd5b80604051602001610acf9190611907565b60405160208183030381529060405280519060200120600060035481548110610b21577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020018190555060036000815480929190610b4290611c94565b919050555050565b6000806000805b6000851115610bdf576001851415610b8c578080610b6e90611c94565b9150508180610b7c90611c94565b9250508181935093505050610be8565b6000600286610b9b9190611d0b565b14610baf578480610bab90611c94565b9550505b8080610bba90611c94565b9150508482610bc99190611b20565b9150600285610bd89190611b76565b9450610b51565b81819350935050505b915091565b8060018190555050565b600060605b60035482101561128257600080839050600185610c199190611ba7565b67ffffffffffffffff811115610c58577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015610c865781602001602082028036833780820191505090505b5092505b600185610c979190611ba7565b82101561120557600080600283610cae9190611d0b565b1415610f4e5786600183610cc29190611b20565b81518110610cf9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151848481518110610d3a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018181525050868281518110610d7f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015187600184610d959190611b20565b81518110610dcc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151604051602001610de59291906118db565b60405160208183030381529060405280519060200120905086600183610e0b9190611b20565b81518110610e42577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151878381518110610e83577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101511115610f495786600183610e9f9190611b20565b81518110610ed6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151878381518110610f17577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151604051602001610f309291906118db565b6040516020818303038152906040528051906020012090505b6111e4565b86600183610f5c9190611ba7565b81518110610f93577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151848481518110610fd4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018181525050868281518110611019577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518760018461102f9190611ba7565b81518110611066577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160405160200161107f9291906118db565b604051602081830303815290604052805190602001209050866001836110a59190611ba7565b815181106110dc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015187838151811061111d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015111156111e357866001836111399190611ba7565b81518110611170577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518783815181106111b1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040516020016111ca9291906118db565b6040516020818303038152906040528051906020012090505b5b6111ef878284611405565b915082806111fc90611c94565b93505050610c8a565b8260046000888781518110611243577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518152602001908152602001600020908051906020019061126c929190611492565b50838061127890611c94565b9450505050610bfc565b50505050565b6000803360405160200161129c91906118a5565b60405160208183030381529060405280519060200120905060005b83518110156113f7578381815181106112f9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518210156113785781848281518110611342577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160405160200161135b9291906118db565b6040516020818303038152906040528051906020012091506113e4565b8381815181106113b1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151826040516020016113cb9291906118db565b6040516020818303038152906040528051906020012091505b80806113ef90611c94565b9150506112b7565b600154821492505050919050565b6000808290505b84518110156114795784818151811061144e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151841415611466578091505061148b565b808061147190611c94565b91505061140c565b50600184516114889190611b20565b90505b9392505050565b8280548282559060005260206000209081019282156114ce579160200282015b828111156114cd5782518255916020019190600101906114b2565b5b5090506114db91906114df565b5090565b5b808211156114f85760008160009055506001016114e0565b5090565b600061150f61150a84611a58565b611a33565b9050808382526020820190508285602086028201111561152e57600080fd5b60005b8581101561155e578161154488826115d0565b845260208401935060208301925050600181019050611531565b5050509392505050565b600061157b61157684611a84565b611a33565b90508281526020810184848401111561159357600080fd5b61159e848285611c21565b509392505050565b600082601f8301126115b757600080fd5b81356115c78482602086016114fc565b91505092915050565b6000813590506115df81611eb1565b92915050565b600082601f8301126115f657600080fd5b8135611606848260208601611568565b91505092915050565b60008135905061161e81611ec8565b92915050565b60006020828403121561163657600080fd5b600082013567ffffffffffffffff81111561165057600080fd5b61165c848285016115a6565b91505092915050565b60006020828403121561167757600080fd5b600082013567ffffffffffffffff81111561169157600080fd5b61169d848285016115e5565b91505092915050565b6000602082840312156116b857600080fd5b60006116c68482850161160f565b91505092915050565b60006116db838361176b565b60208301905092915050565b6116f081611bdb565b82525050565b61170761170282611bdb565b611cdd565b82525050565b600061171882611ac5565b6117228185611af3565b935061172d83611ab5565b8060005b8381101561175e57815161174588826116cf565b975061175083611ae6565b925050600181019050611731565b5085935050505092915050565b61177481611bed565b82525050565b61178381611bed565b82525050565b61179a61179582611bed565b611cef565b82525050565b60006117ab82611ad0565b6117b58185611b04565b93506117c5818560208601611c30565b80840191505092915050565b60006117dc82611adb565b6117e68185611b0f565b93506117f6818560208601611c30565b6117ff81611dc9565b840191505092915050565b6000611817600e83611b0f565b915061182282611de7565b602082019050919050565b600061183a601583611b0f565b915061184582611e10565b602082019050919050565b600061185d601783611b0f565b915061186882611e39565b602082019050919050565b6000611880602c83611b0f565b915061188b82611e62565b604082019050919050565b61189f81611c17565b82525050565b60006118b182846116f6565b60148201915081905092915050565b60006118cc8284611789565b60208201915081905092915050565b60006118e78285611789565b6020820191506118f78284611789565b6020820191508190509392505050565b600061191382846117a0565b915081905092915050565b600060208201905061193360008301846116e7565b92915050565b60006020820190508181036000830152611953818461170d565b905092915050565b6000602082019050611970600083018461177a565b92915050565b6000602082019050818103600083015261199081846117d1565b905092915050565b600060208201905081810360008301526119b18161180a565b9050919050565b600060208201905081810360008301526119d18161182d565b9050919050565b600060208201905081810360008301526119f181611850565b9050919050565b60006020820190508181036000830152611a1181611873565b9050919050565b6000602082019050611a2d6000830184611896565b92915050565b6000611a3d611a4e565b9050611a498282611c63565b919050565b6000604051905090565b600067ffffffffffffffff821115611a7357611a72611d9a565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611a9f57611a9e611d9a565b5b611aa882611dc9565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611b2b82611c17565b9150611b3683611c17565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611b6b57611b6a611d3c565b5b828201905092915050565b6000611b8182611c17565b9150611b8c83611c17565b925082611b9c57611b9b611d6b565b5b828204905092915050565b6000611bb282611c17565b9150611bbd83611c17565b925082821015611bd057611bcf611d3c565b5b828203905092915050565b6000611be682611bf7565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611c4e578082015181840152602081019050611c33565b83811115611c5d576000848401525b50505050565b611c6c82611dc9565b810181811067ffffffffffffffff82111715611c8b57611c8a611d9a565b5b80604052505050565b6000611c9f82611c17565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611cd257611cd1611d3c565b5b600182019050919050565b6000611ce882611cf9565b9050919050565b6000819050919050565b6000611d0482611dda565b9050919050565b6000611d1682611c17565b9150611d2183611c17565b925082611d3157611d30611d6b565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f43616e277420616464206d6f7265000000000000000000000000000000000000600082015250565b7f49742773206e6f742073746172746564207965742e0000000000000000000000600082015250565b7f596f75206172656e2774207768697465206c6973746564000000000000000000600082015250565b7f536f7272792c204f6e6c792041646d696e2063616e206368616e67652074686560008201527f204d65726b6c6520526f6f740000000000000000000000000000000000000000602082015250565b611eba81611bed565b8114611ec557600080fd5b50565b611ed181611c17565b8114611edc57600080fd5b5056fe436f6e67726164756c6174696f6e7321212120596f75277665206265656e207768697465206c697374656421a2646970667358221220f4eaf5daf84e9743ca47441de741e43ab71e3fe82519239485c3957157ab60e664736f6c63430008010033",
}

// NfthiteListABI is the input ABI used to generate the binding from.
// Deprecated: Use NfthiteListMetaData.ABI instead.
var NfthiteListABI = NfthiteListMetaData.ABI

// NfthiteListBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NfthiteListMetaData.Bin instead.
var NfthiteListBin = NfthiteListMetaData.Bin

// DeployNfthiteList deploys a new Ethereum contract, binding an instance of NfthiteList to it.
func DeployNfthiteList(auth *bind.TransactOpts, backend bind.ContractBackend, _max_allowed *big.Int) (common.Address, *types.Transaction, *NfthiteList, error) {
	parsed, err := NfthiteListMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NfthiteListBin), backend, _max_allowed)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NfthiteList{NfthiteListCaller: NfthiteListCaller{contract: contract}, NfthiteListTransactor: NfthiteListTransactor{contract: contract}, NfthiteListFilterer: NfthiteListFilterer{contract: contract}}, nil
}

// NfthiteList is an auto generated Go binding around an Ethereum contract.
type NfthiteList struct {
	NfthiteListCaller     // Read-only binding to the contract
	NfthiteListTransactor // Write-only binding to the contract
	NfthiteListFilterer   // Log filterer for contract events
}

// NfthiteListCaller is an auto generated read-only Go binding around an Ethereum contract.
type NfthiteListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NfthiteListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NfthiteListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NfthiteListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NfthiteListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NfthiteListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NfthiteListSession struct {
	Contract     *NfthiteList      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NfthiteListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NfthiteListCallerSession struct {
	Contract *NfthiteListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NfthiteListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NfthiteListTransactorSession struct {
	Contract     *NfthiteListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NfthiteListRaw is an auto generated low-level Go binding around an Ethereum contract.
type NfthiteListRaw struct {
	Contract *NfthiteList // Generic contract binding to access the raw methods on
}

// NfthiteListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NfthiteListCallerRaw struct {
	Contract *NfthiteListCaller // Generic read-only contract binding to access the raw methods on
}

// NfthiteListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NfthiteListTransactorRaw struct {
	Contract *NfthiteListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNfthiteList creates a new instance of NfthiteList, bound to a specific deployed contract.
func NewNfthiteList(address common.Address, backend bind.ContractBackend) (*NfthiteList, error) {
	contract, err := bindNfthiteList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NfthiteList{NfthiteListCaller: NfthiteListCaller{contract: contract}, NfthiteListTransactor: NfthiteListTransactor{contract: contract}, NfthiteListFilterer: NfthiteListFilterer{contract: contract}}, nil
}

// NewNfthiteListCaller creates a new read-only instance of NfthiteList, bound to a specific deployed contract.
func NewNfthiteListCaller(address common.Address, caller bind.ContractCaller) (*NfthiteListCaller, error) {
	contract, err := bindNfthiteList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NfthiteListCaller{contract: contract}, nil
}

// NewNfthiteListTransactor creates a new write-only instance of NfthiteList, bound to a specific deployed contract.
func NewNfthiteListTransactor(address common.Address, transactor bind.ContractTransactor) (*NfthiteListTransactor, error) {
	contract, err := bindNfthiteList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NfthiteListTransactor{contract: contract}, nil
}

// NewNfthiteListFilterer creates a new log filterer instance of NfthiteList, bound to a specific deployed contract.
func NewNfthiteListFilterer(address common.Address, filterer bind.ContractFilterer) (*NfthiteListFilterer, error) {
	contract, err := bindNfthiteList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NfthiteListFilterer{contract: contract}, nil
}

// bindNfthiteList binds a generic wrapper to an already deployed contract.
func bindNfthiteList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NfthiteListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NfthiteList *NfthiteListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NfthiteList.Contract.NfthiteListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NfthiteList *NfthiteListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NfthiteList.Contract.NfthiteListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NfthiteList *NfthiteListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NfthiteList.Contract.NfthiteListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NfthiteList *NfthiteListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NfthiteList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NfthiteList *NfthiteListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NfthiteList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NfthiteList *NfthiteListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NfthiteList.Contract.contract.Transact(opts, method, params...)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0xa4730acf.
//
// Solidity: function GetMerkleRoot() view returns(bytes32)
func (_NfthiteList *NfthiteListCaller) GetMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NfthiteList.contract.Call(opts, &out, "GetMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMerkleRoot is a free data retrieval call binding the contract method 0xa4730acf.
//
// Solidity: function GetMerkleRoot() view returns(bytes32)
func (_NfthiteList *NfthiteListSession) GetMerkleRoot() ([32]byte, error) {
	return _NfthiteList.Contract.GetMerkleRoot(&_NfthiteList.CallOpts)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0xa4730acf.
//
// Solidity: function GetMerkleRoot() view returns(bytes32)
func (_NfthiteList *NfthiteListCallerSession) GetMerkleRoot() ([32]byte, error) {
	return _NfthiteList.Contract.GetMerkleRoot(&_NfthiteList.CallOpts)
}

// SayHelloToWhiteListed is a free data retrieval call binding the contract method 0x87f8151c.
//
// Solidity: function SayHelloToWhiteListed(bytes32[] _proof) view returns(string)
func (_NfthiteList *NfthiteListCaller) SayHelloToWhiteListed(opts *bind.CallOpts, _proof [][32]byte) (string, error) {
	var out []interface{}
	err := _NfthiteList.contract.Call(opts, &out, "SayHelloToWhiteListed", _proof)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SayHelloToWhiteListed is a free data retrieval call binding the contract method 0x87f8151c.
//
// Solidity: function SayHelloToWhiteListed(bytes32[] _proof) view returns(string)
func (_NfthiteList *NfthiteListSession) SayHelloToWhiteListed(_proof [][32]byte) (string, error) {
	return _NfthiteList.Contract.SayHelloToWhiteListed(&_NfthiteList.CallOpts, _proof)
}

// SayHelloToWhiteListed is a free data retrieval call binding the contract method 0x87f8151c.
//
// Solidity: function SayHelloToWhiteListed(bytes32[] _proof) view returns(string)
func (_NfthiteList *NfthiteListCallerSession) SayHelloToWhiteListed(_proof [][32]byte) (string, error) {
	return _NfthiteList.Contract.SayHelloToWhiteListed(&_NfthiteList.CallOpts, _proof)
}

// AddedItems is a free data retrieval call binding the contract method 0x7415ad99.
//
// Solidity: function addedItems() view returns(uint256)
func (_NfthiteList *NfthiteListCaller) AddedItems(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NfthiteList.contract.Call(opts, &out, "addedItems")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddedItems is a free data retrieval call binding the contract method 0x7415ad99.
//
// Solidity: function addedItems() view returns(uint256)
func (_NfthiteList *NfthiteListSession) AddedItems() (*big.Int, error) {
	return _NfthiteList.Contract.AddedItems(&_NfthiteList.CallOpts)
}

// AddedItems is a free data retrieval call binding the contract method 0x7415ad99.
//
// Solidity: function addedItems() view returns(uint256)
func (_NfthiteList *NfthiteListCallerSession) AddedItems() (*big.Int, error) {
	return _NfthiteList.Contract.AddedItems(&_NfthiteList.CallOpts)
}

// GetProof is a free data retrieval call binding the contract method 0x5c5d625e.
//
// Solidity: function getProof() view returns(bytes32[])
func (_NfthiteList *NfthiteListCaller) GetProof(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _NfthiteList.contract.Call(opts, &out, "getProof")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x5c5d625e.
//
// Solidity: function getProof() view returns(bytes32[])
func (_NfthiteList *NfthiteListSession) GetProof() ([][32]byte, error) {
	return _NfthiteList.Contract.GetProof(&_NfthiteList.CallOpts)
}

// GetProof is a free data retrieval call binding the contract method 0x5c5d625e.
//
// Solidity: function getProof() view returns(bytes32[])
func (_NfthiteList *NfthiteListCallerSession) GetProof() ([][32]byte, error) {
	return _NfthiteList.Contract.GetProof(&_NfthiteList.CallOpts)
}

// ListingRequests is a free data retrieval call binding the contract method 0xa04e63bb.
//
// Solidity: function listingRequests(uint256 ) view returns(bytes32)
func (_NfthiteList *NfthiteListCaller) ListingRequests(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _NfthiteList.contract.Call(opts, &out, "listingRequests", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ListingRequests is a free data retrieval call binding the contract method 0xa04e63bb.
//
// Solidity: function listingRequests(uint256 ) view returns(bytes32)
func (_NfthiteList *NfthiteListSession) ListingRequests(arg0 *big.Int) ([32]byte, error) {
	return _NfthiteList.Contract.ListingRequests(&_NfthiteList.CallOpts, arg0)
}

// ListingRequests is a free data retrieval call binding the contract method 0xa04e63bb.
//
// Solidity: function listingRequests(uint256 ) view returns(bytes32)
func (_NfthiteList *NfthiteListCallerSession) ListingRequests(arg0 *big.Int) ([32]byte, error) {
	return _NfthiteList.Contract.ListingRequests(&_NfthiteList.CallOpts, arg0)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_NfthiteList *NfthiteListCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NfthiteList.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_NfthiteList *NfthiteListSession) MerkleRoot() ([32]byte, error) {
	return _NfthiteList.Contract.MerkleRoot(&_NfthiteList.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_NfthiteList *NfthiteListCallerSession) MerkleRoot() ([32]byte, error) {
	return _NfthiteList.Contract.MerkleRoot(&_NfthiteList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NfthiteList *NfthiteListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NfthiteList.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NfthiteList *NfthiteListSession) Owner() (common.Address, error) {
	return _NfthiteList.Contract.Owner(&_NfthiteList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NfthiteList *NfthiteListCallerSession) Owner() (common.Address, error) {
	return _NfthiteList.Contract.Owner(&_NfthiteList.CallOpts)
}

// AddAddressToWhiteList is a paid mutator transaction binding the contract method 0xe3ad6acf.
//
// Solidity: function AddAddressToWhiteList(bytes addr) returns()
func (_NfthiteList *NfthiteListTransactor) AddAddressToWhiteList(opts *bind.TransactOpts, addr []byte) (*types.Transaction, error) {
	return _NfthiteList.contract.Transact(opts, "AddAddressToWhiteList", addr)
}

// AddAddressToWhiteList is a paid mutator transaction binding the contract method 0xe3ad6acf.
//
// Solidity: function AddAddressToWhiteList(bytes addr) returns()
func (_NfthiteList *NfthiteListSession) AddAddressToWhiteList(addr []byte) (*types.Transaction, error) {
	return _NfthiteList.Contract.AddAddressToWhiteList(&_NfthiteList.TransactOpts, addr)
}

// AddAddressToWhiteList is a paid mutator transaction binding the contract method 0xe3ad6acf.
//
// Solidity: function AddAddressToWhiteList(bytes addr) returns()
func (_NfthiteList *NfthiteListTransactorSession) AddAddressToWhiteList(addr []byte) (*types.Transaction, error) {
	return _NfthiteList.Contract.AddAddressToWhiteList(&_NfthiteList.TransactOpts, addr)
}

// BuildMerkleTree is a paid mutator transaction binding the contract method 0x7d3e9a78.
//
// Solidity: function BuildMerkleTree() returns()
func (_NfthiteList *NfthiteListTransactor) BuildMerkleTree(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NfthiteList.contract.Transact(opts, "BuildMerkleTree")
}

// BuildMerkleTree is a paid mutator transaction binding the contract method 0x7d3e9a78.
//
// Solidity: function BuildMerkleTree() returns()
func (_NfthiteList *NfthiteListSession) BuildMerkleTree() (*types.Transaction, error) {
	return _NfthiteList.Contract.BuildMerkleTree(&_NfthiteList.TransactOpts)
}

// BuildMerkleTree is a paid mutator transaction binding the contract method 0x7d3e9a78.
//
// Solidity: function BuildMerkleTree() returns()
func (_NfthiteList *NfthiteListTransactorSession) BuildMerkleTree() (*types.Transaction, error) {
	return _NfthiteList.Contract.BuildMerkleTree(&_NfthiteList.TransactOpts)
}
