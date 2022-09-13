// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const GasPriceOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":0,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":0,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"spacer0\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":0,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"spacer1\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":0,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint256\"},{\"astId\":0,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_uint256\"},{\"astId\":0,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"decimals\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var GasPriceOracleStorageLayout = new(solc.StorageLayout)

var GasPriceOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106100f55760003560e01c80637046559711610097578063de26c4a111610066578063de26c4a1146101c0578063f2fde38b146101d3578063f45e65d8146101e6578063fe173b971461016457600080fd5b8063704655971461016a578063715018a61461017d5780638c8885c8146101855780638da5cb5b1461019857600080fd5b806349948e0e116100d357806349948e0e14610134578063519b4bd31461014757806354fd4d501461014f5780636ef25c3a1461016457600080fd5b80630c18c162146100fa578063313ce567146101165780633577afc51461011f575b600080fd5b61010360035481565b6040519081526020015b60405180910390f35b61010360055481565b61013261012d3660046107e8565b6101ef565b005b610103610142366004610830565b610233565b610103610293565b61015761031d565b60405161010d919061092f565b48610103565b6101326101783660046107e8565b6103c0565b6101326103fd565b6101326101933660046107e8565b610411565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010d565b6101036101ce366004610830565b61044e565b6101326101e1366004610980565b6104f9565b61010360045481565b6101f76105b5565b60038190556040518181527f32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4906020015b60405180910390a150565b60008061023f8361044e565b9050600061024b610293565b61025590836109ec565b90506000600554600a6102689190610b4b565b905060006004548361027a91906109ec565b905060006102888383610b86565b979650505050505050565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16635cf249696040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103189190610b9a565b905090565b60606103487f0000000000000000000000000000000000000000000000000000000000000000610636565b6103717f0000000000000000000000000000000000000000000000000000000000000000610636565b61039a7f0000000000000000000000000000000000000000000000000000000000000000610636565b6040516020016103ac93929190610bb3565b604051602081830303815290604052905090565b6103c86105b5565b60048190556040518181527f3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a90602001610228565b6104056105b5565b61040f6000610773565b565b6104196105b5565b60058190556040518181527fd68112a8707e326d08be3656b528c1bcc5bbbfc47f4177e2179b14d8640838c190602001610228565b80516000908190815b818110156104d15784818151811061047157610471610c29565b01602001517fff00000000000000000000000000000000000000000000000000000000000000166000036104b1576104aa600484610c58565b92506104bf565b6104bc601084610c58565b92505b806104c981610c70565b915050610457565b506000600354836104e29190610c58565b90506104f081610440610c58565b95945050505050565b6105016105b5565b73ffffffffffffffffffffffffffffffffffffffff81166105a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6105b281610773565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461040f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105a0565b60608160000361067957505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156106a3578061068d81610c70565b915061069c9050600a83610b86565b915061067d565b60008167ffffffffffffffff8111156106be576106be610801565b6040519080825280601f01601f1916602001820160405280156106e8576020820181803683370190505b5090505b841561076b576106fd600183610ca8565b915061070a600a86610cbf565b610715906030610c58565b60f81b81838151811061072a5761072a610c29565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610764600a86610b86565b94506106ec565b949350505050565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156107fa57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561084257600080fd5b813567ffffffffffffffff8082111561085a57600080fd5b818401915084601f83011261086e57600080fd5b81358181111561088057610880610801565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156108c6576108c6610801565b816040528281528760208487010111156108df57600080fd5b826020860160208301376000928101602001929092525095945050505050565b60005b8381101561091a578181015183820152602001610902565b83811115610929576000848401525b50505050565b602081526000825180602084015261094e8160408501602087016108ff565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60006020828403121561099257600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146109b657600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610a2457610a246109bd565b500290565b600181815b80851115610a8257817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610a6857610a686109bd565b80851615610a7557918102915b93841c9390800290610a2e565b509250929050565b600082610a9957506001610b45565b81610aa657506000610b45565b8160018114610abc5760028114610ac657610ae2565b6001915050610b45565b60ff841115610ad757610ad76109bd565b50506001821b610b45565b5060208310610133831016604e8410600b8410161715610b05575081810a610b45565b610b0f8383610a29565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610b4157610b416109bd565b0290505b92915050565b60006109b68383610a8a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610b9557610b95610b57565b500490565b600060208284031215610bac57600080fd5b5051919050565b60008451610bc58184602089016108ff565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610c01816001850160208a016108ff565b60019201918201528351610c1c8160028401602088016108ff565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008219821115610c6b57610c6b6109bd565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610ca157610ca16109bd565b5060010190565b600082821015610cba57610cba6109bd565b500390565b600082610cce57610cce610b57565b50069056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(GasPriceOracleStorageLayoutJSON), GasPriceOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["GasPriceOracle"] = GasPriceOracleStorageLayout
	deployedBytecodes["GasPriceOracle"] = GasPriceOracleDeployedBin
}