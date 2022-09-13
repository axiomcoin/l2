// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer0\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_paused\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_status\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer1\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_uint256\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer2\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_uint256\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"otherMessenger\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_address\"},{\"astId\":0,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"receivedMessages\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_mapping(t_bytes32,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2CrossDomainMessengerDeployedBin = "0x6080604052600436106101755760003560e01c80638456cb59116100cb578063c4d66de81161007f578063ecc7042811610059578063ecc70428146103ee578063f2fde38b14610453578063f69f81511461047357600080fd5b8063c4d66de81461038e578063d764ad0b146103ae578063db505d80146103c157600080fd5b8063a7119869116100b0578063a711986914610313578063b1b1b2091461033e578063b28ade251461036e57600080fd5b80638456cb59146102d35780638da5cb5b146102e857600080fd5b80633f827a5a1161012d5780636e296e45116101075780636e296e451461026d578063715018a6146102a75780637dea7cc3146102bc57600080fd5b80633f827a5a146101ff57806354fd4d50146102275780635c975abb1461024957600080fd5b80632828d7e81161015e5780632828d7e8146101bf5780633dbb202b146101d55780633f4ba83a146101ea57600080fd5b8063028f85f71461017a5780630c568498146101a9575b600080fd5b34801561018657600080fd5b5061018f601081565b60405163ffffffff90911681526020015b60405180910390f35b3480156101b557600080fd5b5061018f6103e881565b3480156101cb57600080fd5b5061018f6103f881565b6101e86101e3366004611d50565b6104a3565b005b3480156101f657600080fd5b506101e8610706565b34801561020b57600080fd5b50610214600181565b60405161ffff90911681526020016101a0565b34801561023357600080fd5b5061023c610718565b6040516101a09190611e2f565b34801561025557600080fd5b5060655460ff165b60405190151581526020016101a0565b34801561027957600080fd5b506102826107bb565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a0565b3480156102b357600080fd5b506101e86108a7565b3480156102c857600080fd5b5061018f62030d4081565b3480156102df57600080fd5b506101e86108b9565b3480156102f457600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610282565b34801561031f57600080fd5b5060ce5473ffffffffffffffffffffffffffffffffffffffff16610282565b34801561034a57600080fd5b5061025d610359366004611e49565b60cb6020526000908152604090205460ff1681565b34801561037a57600080fd5b5061018f610389366004611e62565b6108c9565b34801561039a57600080fd5b506101e86103a9366004611eb6565b61090f565b6101e86103bc366004611ed1565b610b0e565b3480156103cd57600080fd5b5060ce546102829073ffffffffffffffffffffffffffffffffffffffff1681565b3480156103fa57600080fd5b5061044560cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016101a0565b34801561045f57600080fd5b506101e861046e366004611eb6565b611259565b34801561047f57600080fd5b5061025d61048e366004611e49565b60cf6020526000908152604090205460ff1681565b60ce546105db9073ffffffffffffffffffffffffffffffffffffffff166104cb8585856108c9565b63ffffffff16347fd764ad0b0000000000000000000000000000000000000000000000000000000061053d60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016105599796959493929190611f9c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261132c565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561066060cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610672959493929190611ffb565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b61070e6113ba565b61071661143b565b565b60606107437f00000000000000000000000000000000000000000000000000000000000000006114b8565b61076c7f00000000000000000000000000000000000000000000000000000000000000006114b8565b6107957f00000000000000000000000000000000000000000000000000000000000000006114b8565b6040516020016107a793929190612049565b604051602081830303815290604052905090565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21530161088a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6108af6113ba565b61071660006115ed565b6108c16113ba565b610716611664565b600062030d406108da6010856120ee565b6103e86108e96103f8866120ee565b6108f39190612149565b6108fd919061216c565b610907919061216c565b949350505050565b6000547501000000000000000000000000000000000000000000900460ff161580801561095a575060005460017401000000000000000000000000000000000000000090910460ff16105b8061098c5750303b15801561098c575060005474010000000000000000000000000000000000000000900460ff166001145b610a18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610881565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790558015610a9e57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b610aa7826116bf565b8015610b0a57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b600260975403610b7a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610881565b6002609755610b876117da565b60f087901c60018114610c42576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605560248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2031206d657373616765732061726520737570706f72746564206166746560648201527f722074686520426564726f636b20757067726164650000000000000000000000608482015260a401610881565b6000610c88898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061184792505050565b9050610cd160ce54337fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef0173ffffffffffffffffffffffffffffffffffffffff90811691161490565b15610d6a57853414610d65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f43726f7373446f6d61696e4d657373656e6765723a206d69736d61746368656460448201527f206d6573736167652076616c75650000000000000000000000000000000000006064820152608401610881565b610ebc565b3415610e1e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a401610881565b600081815260cf602052604090205460ff16610ebc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c61796564000000000000000000000000000000006064820152608401610881565b610ec58761186a565b15610f78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a401610881565b600081815260cb602052604090205460ff1615611017576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c61796564000000000000000000006064820152608401610881565b61102361afc886612194565b5a10156110b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a20696e737566666963696560448201527f6e742067617320746f2072656c6179206d6573736167650000000000000000006064820152608401610881565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a16179055600061114e8861110661138861afc86121ac565b5a61111191906121ac565b8988888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506118bf92505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905590508015156001036111e957600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a2611248565b600082815260cf602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a25b505060016097555050505050505050565b6112616113ba565b73ffffffffffffffffffffffffffffffffffffffff8116611304576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610881565b61130d816115ed565b50565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6040517fc2b3e5ac0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000009063c2b3e5ac908490611382908890889087906004016121c3565b6000604051808303818588803b15801561139b57600080fd5b505af11580156113af573d6000803e3d6000fd5b505050505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610716576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610881565b6114436118d9565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6060816000036114fb57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611525578061150f8161220b565b915061151e9050600a83612243565b91506114ff565b60008167ffffffffffffffff81111561154057611540612257565b6040519080825280601f01601f19166020018201604052801561156a576020820181803683370190505b5090505b84156109075761157f6001836121ac565b915061158c600a86612286565b611597906030612194565b60f81b8183815181106115ac576115ac61229a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506115e6600a86612243565b945061156e565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b61166c6117da565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861148e3390565b6000547501000000000000000000000000000000000000000000900460ff1661176a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610881565b60cc805461dead7fffffffffffffffffffffffff00000000000000000000000000000000000000009182161790915560ce805490911673ffffffffffffffffffffffffffffffffffffffff83161790556117c2611945565b6117ca6119f0565b6117d2611aa4565b61130d611b79565b60655460ff1615610716576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610881565b6000611857878787878787611c2b565b8051906020012090509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff82163014806118b9575073ffffffffffffffffffffffffffffffffffffffff8216734200000000000000000000000000000000000000145b92915050565b600080600080845160208601878a8af19695505050505050565b60655460ff16610716576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610881565b6000547501000000000000000000000000000000000000000000900460ff16610716576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610881565b6000547501000000000000000000000000000000000000000000900460ff16611a9b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610881565b610716336115ed565b6000547501000000000000000000000000000000000000000000900460ff16611b4f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610881565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6000547501000000000000000000000000000000000000000000900460ff16611c24576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610881565b6001609755565b6060868686868686604051602401611c48969594939291906122c9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611cee57600080fd5b919050565b60008083601f840112611d0557600080fd5b50813567ffffffffffffffff811115611d1d57600080fd5b602083019150836020828501011115611d3557600080fd5b9250929050565b803563ffffffff81168114611cee57600080fd5b60008060008060608587031215611d6657600080fd5b611d6f85611cca565b9350602085013567ffffffffffffffff811115611d8b57600080fd5b611d9787828801611cf3565b9094509250611daa905060408601611d3c565b905092959194509250565b60005b83811015611dd0578181015183820152602001611db8565b83811115611ddf576000848401525b50505050565b60008151808452611dfd816020860160208601611db5565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611e426020830184611de5565b9392505050565b600060208284031215611e5b57600080fd5b5035919050565b600080600060408486031215611e7757600080fd5b833567ffffffffffffffff811115611e8e57600080fd5b611e9a86828701611cf3565b9094509250611ead905060208501611d3c565b90509250925092565b600060208284031215611ec857600080fd5b611e4282611cca565b600080600080600080600060c0888a031215611eec57600080fd5b87359650611efc60208901611cca565b9550611f0a60408901611cca565b9450606088013593506080880135925060a088013567ffffffffffffffff811115611f3457600080fd5b611f408a828b01611cf3565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611fee60c083018486611f53565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815260806020820152600061202b608083018688611f53565b905083604083015263ffffffff831660608301529695505050505050565b6000845161205b818460208901611db5565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551612097816001850160208a01611db5565b600192019182015283516120b2816002840160208801611db5565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600063ffffffff80831681851681830481118215151615612111576121116120bf565b02949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600063ffffffff808416806121605761216061211a565b92169190910492915050565b600063ffffffff80831681851680830382111561218b5761218b6120bf565b01949350505050565b600082198211156121a7576121a76120bf565b500190565b6000828210156121be576121be6120bf565b500390565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff831660208201526060604082015260006122026060830184611de5565b95945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361223c5761223c6120bf565b5060010190565b6000826122525761225261211a565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000826122955761229561211a565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a083015261231460c0830184611de5565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2CrossDomainMessengerStorageLayoutJSON), L2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2CrossDomainMessenger"] = L2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2CrossDomainMessenger"] = L2CrossDomainMessengerDeployedBin
}