// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier_proxy

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type CommonAddressAndWeight struct {
	Addr   common.Address
	Weight uint64
}

var VerifierProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"accessController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessForbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadVerification\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"ConfigDigestAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeManagerInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VerifierInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"VerifierNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAccessController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAccessController\",\"type\":\"address\"}],\"name\":\"AccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"FeeManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"VerifierInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldConfigDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newConfigDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"VerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"VerifierUnset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"initializeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_accessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_feeManager\",\"outputs\":[{\"internalType\":\"contractIVerifierFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"accessController\",\"type\":\"address\"}],\"name\":\"setAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVerifierFeeManager\",\"name\":\"feeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currentConfigDigest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newConfigDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structCommon.AddressAndWeight[]\",\"name\":\"addressesAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"unsetVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"}],\"name\":\"verifyBulk\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"verifiedReports\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001c6138038062001c61833981016040819052620000349162000193565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e8565b5050600480546001600160a01b0319166001600160a01b03939093169290921790915550620001c5565b336001600160a01b03821603620001425760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001a657600080fd5b81516001600160a01b0381168114620001be57600080fd5b9392505050565b611a8c80620001d56000396000f3fe6080604052600436106100dd5760003560e01c80638da5cb5b1161007f578063b011b24711610059578063b011b2471461028e578063eeb7b248146102ae578063f08391d8146102f1578063f2fde38b1461031157600080fd5b80638da5cb5b146102235780638e760afe1461024e57806394ba28461461026157600080fd5b80636e914094116100bb5780636e914094146101ae57806379ba5097146101ce57806383490cd7146101e35780638c2a4d531461020357600080fd5b8063181f5a77146100e257806338416b5b1461013a578063472d35b91461018c575b600080fd5b3480156100ee57600080fd5b5060408051808201909152601381527f566572696669657250726f787920312e312e300000000000000000000000000060208201525b60405161013191906113b2565b60405180910390f35b34801561014657600080fd5b506005546101679073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610131565b34801561019857600080fd5b506101ac6101a73660046113ee565b610331565b005b3480156101ba57600080fd5b506101ac6101c936600461140b565b6105a9565b3480156101da57600080fd5b506101ac61069a565b6101f66101f1366004611424565b610797565b6040516101319190611499565b34801561020f57600080fd5b506101ac61021e3660046113ee565b6109f5565b34801561022f57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610167565b61012461025c366004611519565b610c26565b34801561026d57600080fd5b506004546101679073ffffffffffffffffffffffffffffffffffffffff1681565b34801561029a57600080fd5b506101ac6102a9366004611579565b610dda565b3480156102ba57600080fd5b506101676102c936600461140b565b60009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b3480156102fd57600080fd5b506101ac61030c3660046113ee565b611000565b34801561031d57600080fd5b506101ac61032c3660046113ee565b611087565b61033961109b565b73ffffffffffffffffffffffffffffffffffffffff8116610386576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527ff1387e1600000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff8216906301ffc9a790602401602060405180830381865afa158015610410573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043491906115fc565b15806104eb57506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f40d7f78e00000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff8216906301ffc9a790602401602060405180830381865afa1580156104c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e991906115fc565b155b15610522576040517f8238941900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f04628abcaa6b1674651352125cb94b65b289145bc2bc4d67720bb7d966372f0391015b60405180910390a15050565b6105b161109b565b60008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1680610615576040517fb151802b000000000000000000000000000000000000000000000000000000008152600481018390526024015b60405180910390fd5b6000828152600360205260409081902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055517f11dc15c4b8ac2b183166cc8427e5385a5ece8308217a4217338c6a7614845c4c9061059d908490849091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b60015473ffffffffffffffffffffffffffffffffffffffff16331461071b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161060c565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60045460609073ffffffffffffffffffffffffffffffffffffffff16801580159061085757506040517f6b14daf800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690636b14daf8906108149033906000903690600401611667565b602060405180830381865afa158015610831573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085591906115fc565b155b1561088e576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60055473ffffffffffffffffffffffffffffffffffffffff168015610938576040517f40d7f78e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216906340d7f78e90349061090590899089903390600401611697565b6000604051808303818588803b15801561091e57600080fd5b505af1158015610932573d6000803e3d6000fd5b50505050505b8367ffffffffffffffff8111156109515761095161178c565b60405190808252806020026020018201604052801561098457816020015b606081526020019060019003908161096f5790505b50925060005b848110156109ec576109be8686838181106109a7576109a76117bb565b90506020028101906109b991906117ea565b61111e565b8482815181106109d0576109d06117bb565b6020026020010181905250806109e590611856565b905061098a565b50505092915050565b6109fd61109b565b8073ffffffffffffffffffffffffffffffffffffffff8116610a4b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f3d3ac1b500000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff8216906301ffc9a790602401602060405180830381865afa158015610ad5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af991906115fc565b610b2f576040517f75b0527a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660009081526002602052604090205460ff1615610ba7576040517f4e01ccfd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161060c565b73ffffffffffffffffffffffffffffffffffffffff821660008181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519182527f1f2cd7c97f4d801b5efe26cc409617c1fd6c5ef786e79aacb90af40923e4e8e9910161059d565b60045460609073ffffffffffffffffffffffffffffffffffffffff168015801590610ce657506040517f6b14daf800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690636b14daf890610ca39033906000903690600401611667565b602060405180830381865afa158015610cc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce491906115fc565b155b15610d1d576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60055473ffffffffffffffffffffffffffffffffffffffff168015610dc7576040517ff1387e1600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82169063f1387e16903490610d94908990899033906004016118b5565b6000604051808303818588803b158015610dad57600080fd5b505af1158015610dc1573d6000803e3d6000fd5b50505050505b610dd1858561111e565b95945050505050565b600083815260036020526040902054839073ffffffffffffffffffffffffffffffffffffffff168015610e58576040517f375d1fe60000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff8216602482015260440161060c565b3360009081526002602052604090205460ff16610ea1576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600085815260036020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016331790558215610fb95760055473ffffffffffffffffffffffffffffffffffffffff16610f2c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005546040517ff65df96200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063f65df96290610f86908890889088906004016118ef565b600060405180830381600087803b158015610fa057600080fd5b505af1158015610fb4573d6000803e3d6000fd5b505050505b6040805187815260208101879052338183015290517fbeb513e532542a562ac35699e7cd9ae7d198dcd3eee15bada6c857d28ceaddcf9181900360600190a1505050505050565b61100861109b565b6004805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f953e92b1a6442e9c3242531154a3f6f6eb00b4e9c719ba8118fa6235e4ce89b6910161059d565b61108f61109b565b6110988161124f565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461111c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161060c565b565b6060600061112c8385611978565b60008181526003602052604090205490915073ffffffffffffffffffffffffffffffffffffffff168061118e576040517fb151802b0000000000000000000000000000000000000000000000000000000081526004810183905260240161060c565b6040517f3d3ac1b500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690633d3ac1b5906111e4908890889033906004016118b5565b6000604051808303816000875af1158015611203573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610dd191908101906119b4565b92915050565b3373ffffffffffffffffffffffffffffffffffffffff8216036112ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161060c565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60005b8381101561135f578181015183820152602001611347565b50506000910152565b60008151808452611380816020860160208601611344565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006113c56020830184611368565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461109857600080fd5b60006020828403121561140057600080fd5b81356113c5816113cc565b60006020828403121561141d57600080fd5b5035919050565b6000806020838503121561143757600080fd5b823567ffffffffffffffff8082111561144f57600080fd5b818501915085601f83011261146357600080fd5b81358181111561147257600080fd5b8660208260051b850101111561148757600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561150c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08886030184526114fa858351611368565b945092850192908501906001016114c0565b5092979650505050505050565b6000806020838503121561152c57600080fd5b823567ffffffffffffffff8082111561154457600080fd5b818501915085601f83011261155857600080fd5b81358181111561156757600080fd5b86602082850101111561148757600080fd5b6000806000806060858703121561158f57600080fd5b8435935060208501359250604085013567ffffffffffffffff808211156115b557600080fd5b818701915087601f8301126115c957600080fd5b8135818111156115d857600080fd5b8860208260061b85010111156115ed57600080fd5b95989497505060200194505050565b60006020828403121561160e57600080fd5b815180151581146113c557600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff84168152604060208201526000610dd160408301848661161e565b6040808252810183905260006060600585901b8301810190830186835b87811015611763577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa086850301835281357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18a360301811261171557600080fd5b8901602081810191359067ffffffffffffffff82111561173457600080fd5b81360383131561174357600080fd5b61174e87838561161e565b965094850194939093019250506001016116b4565b50505073ffffffffffffffffffffffffffffffffffffffff841660208401529050949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261181f57600080fd5b83018035915067ffffffffffffffff82111561183a57600080fd5b60200191503681900382131561184f57600080fd5b9250929050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118ae577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b6040815260006118c960408301858761161e565b905073ffffffffffffffffffffffffffffffffffffffff83166020830152949350505050565b838152604060208083018290528282018490526000919085906060850184805b88811015611969578435611922816113cc565b73ffffffffffffffffffffffffffffffffffffffff1683528484013567ffffffffffffffff8116808214611954578384fd5b8486015250938501939185019160010161190f565b50909998505050505050505050565b80356020831015611249577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602084900360031b1b1692915050565b6000602082840312156119c657600080fd5b815167ffffffffffffffff808211156119de57600080fd5b818401915084601f8301126119f257600080fd5b815181811115611a0457611a0461178c565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611a4a57611a4a61178c565b81604052828152876020848701011115611a6357600080fd5b611a74836020830160208801611344565b97965050505050505056fea164736f6c6343000810000a",
}

var VerifierProxyABI = VerifierProxyMetaData.ABI

var VerifierProxyBin = VerifierProxyMetaData.Bin

func DeployVerifierProxy(auth *bind.TransactOpts, backend bind.ContractBackend, accessController common.Address) (common.Address, *types.Transaction, *VerifierProxy, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierProxyBin), backend, accessController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierProxy{VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

type VerifierProxy struct {
	address common.Address
	abi     abi.ABI
	VerifierProxyCaller
	VerifierProxyTransactor
	VerifierProxyFilterer
}

type VerifierProxyCaller struct {
	contract *bind.BoundContract
}

type VerifierProxyTransactor struct {
	contract *bind.BoundContract
}

type VerifierProxyFilterer struct {
	contract *bind.BoundContract
}

type VerifierProxySession struct {
	Contract     *VerifierProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierProxyCallerSession struct {
	Contract *VerifierProxyCaller
	CallOpts bind.CallOpts
}

type VerifierProxyTransactorSession struct {
	Contract     *VerifierProxyTransactor
	TransactOpts bind.TransactOpts
}

type VerifierProxyRaw struct {
	Contract *VerifierProxy
}

type VerifierProxyCallerRaw struct {
	Contract *VerifierProxyCaller
}

type VerifierProxyTransactorRaw struct {
	Contract *VerifierProxyTransactor
}

func NewVerifierProxy(address common.Address, backend bind.ContractBackend) (*VerifierProxy, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifierProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierProxy{address: address, abi: abi, VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

func NewVerifierProxyCaller(address common.Address, caller bind.ContractCaller) (*VerifierProxyCaller, error) {
	contract, err := bindVerifierProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyCaller{contract: contract}, nil
}

func NewVerifierProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierProxyTransactor, error) {
	contract, err := bindVerifierProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyTransactor{contract: contract}, nil
}

func NewVerifierProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierProxyFilterer, error) {
	contract, err := bindVerifierProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyFilterer{contract: contract}, nil
}

func bindVerifierProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VerifierProxy *VerifierProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.VerifierProxyCaller.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCaller) GetVerifier(opts *bind.CallOpts, configDigest [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getVerifier", configDigest)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetVerifier(configDigest [32]byte) (common.Address, error) {
	return _VerifierProxy.Contract.GetVerifier(&_VerifierProxy.CallOpts, configDigest)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetVerifier(configDigest [32]byte) (common.Address, error) {
	return _VerifierProxy.Contract.GetVerifier(&_VerifierProxy.CallOpts, configDigest)
}

func (_VerifierProxy *VerifierProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) SAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "s_accessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) SAccessController() (common.Address, error) {
	return _VerifierProxy.Contract.SAccessController(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) SAccessController() (common.Address, error) {
	return _VerifierProxy.Contract.SAccessController(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) SFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "s_feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) SFeeManager() (common.Address, error) {
	return _VerifierProxy.Contract.SFeeManager(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) SFeeManager() (common.Address, error) {
	return _VerifierProxy.Contract.SFeeManager(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) TypeAndVersion() (string, error) {
	return _VerifierProxy.Contract.TypeAndVersion(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) TypeAndVersion() (string, error) {
	return _VerifierProxy.Contract.TypeAndVersion(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "acceptOwnership")
}

func (_VerifierProxy *VerifierProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactor) InitializeVerifier(opts *bind.TransactOpts, verifierAddress common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "initializeVerifier", verifierAddress)
}

func (_VerifierProxy *VerifierProxySession) InitializeVerifier(verifierAddress common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.InitializeVerifier(&_VerifierProxy.TransactOpts, verifierAddress)
}

func (_VerifierProxy *VerifierProxyTransactorSession) InitializeVerifier(verifierAddress common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.InitializeVerifier(&_VerifierProxy.TransactOpts, verifierAddress)
}

func (_VerifierProxy *VerifierProxyTransactor) SetAccessController(opts *bind.TransactOpts, accessController common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "setAccessController", accessController)
}

func (_VerifierProxy *VerifierProxySession) SetAccessController(accessController common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetAccessController(&_VerifierProxy.TransactOpts, accessController)
}

func (_VerifierProxy *VerifierProxyTransactorSession) SetAccessController(accessController common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetAccessController(&_VerifierProxy.TransactOpts, accessController)
}

func (_VerifierProxy *VerifierProxyTransactor) SetFeeManager(opts *bind.TransactOpts, feeManager common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "setFeeManager", feeManager)
}

func (_VerifierProxy *VerifierProxySession) SetFeeManager(feeManager common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetFeeManager(&_VerifierProxy.TransactOpts, feeManager)
}

func (_VerifierProxy *VerifierProxyTransactorSession) SetFeeManager(feeManager common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetFeeManager(&_VerifierProxy.TransactOpts, feeManager)
}

func (_VerifierProxy *VerifierProxyTransactor) SetVerifier(opts *bind.TransactOpts, currentConfigDigest [32]byte, newConfigDigest [32]byte, addressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "setVerifier", currentConfigDigest, newConfigDigest, addressesAndWeights)
}

func (_VerifierProxy *VerifierProxySession) SetVerifier(currentConfigDigest [32]byte, newConfigDigest [32]byte, addressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetVerifier(&_VerifierProxy.TransactOpts, currentConfigDigest, newConfigDigest, addressesAndWeights)
}

func (_VerifierProxy *VerifierProxyTransactorSession) SetVerifier(currentConfigDigest [32]byte, newConfigDigest [32]byte, addressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetVerifier(&_VerifierProxy.TransactOpts, currentConfigDigest, newConfigDigest, addressesAndWeights)
}

func (_VerifierProxy *VerifierProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_VerifierProxy *VerifierProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactor) UnsetVerifier(opts *bind.TransactOpts, configDigest [32]byte) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "unsetVerifier", configDigest)
}

func (_VerifierProxy *VerifierProxySession) UnsetVerifier(configDigest [32]byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.UnsetVerifier(&_VerifierProxy.TransactOpts, configDigest)
}

func (_VerifierProxy *VerifierProxyTransactorSession) UnsetVerifier(configDigest [32]byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.UnsetVerifier(&_VerifierProxy.TransactOpts, configDigest)
}

func (_VerifierProxy *VerifierProxyTransactor) Verify(opts *bind.TransactOpts, payload []byte) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "verify", payload)
}

func (_VerifierProxy *VerifierProxySession) Verify(payload []byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.Verify(&_VerifierProxy.TransactOpts, payload)
}

func (_VerifierProxy *VerifierProxyTransactorSession) Verify(payload []byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.Verify(&_VerifierProxy.TransactOpts, payload)
}

func (_VerifierProxy *VerifierProxyTransactor) VerifyBulk(opts *bind.TransactOpts, payloads [][]byte) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "verifyBulk", payloads)
}

func (_VerifierProxy *VerifierProxySession) VerifyBulk(payloads [][]byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifyBulk(&_VerifierProxy.TransactOpts, payloads)
}

func (_VerifierProxy *VerifierProxyTransactorSession) VerifyBulk(payloads [][]byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifyBulk(&_VerifierProxy.TransactOpts, payloads)
}

type VerifierProxyAccessControllerSetIterator struct {
	Event *VerifierProxyAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyAccessControllerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierProxyAccessControllerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierProxyAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyAccessControllerSet struct {
	OldAccessController common.Address
	NewAccessController common.Address
	Raw                 types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterAccessControllerSet(opts *bind.FilterOpts) (*VerifierProxyAccessControllerSetIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "AccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyAccessControllerSetIterator{contract: _VerifierProxy.contract, event: "AccessControllerSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchAccessControllerSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "AccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyAccessControllerSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "AccessControllerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VerifierProxy *VerifierProxyFilterer) ParseAccessControllerSet(log types.Log) (*VerifierProxyAccessControllerSet, error) {
	event := new(VerifierProxyAccessControllerSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "AccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyFeeManagerSetIterator struct {
	Event *VerifierProxyFeeManagerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyFeeManagerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyFeeManagerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierProxyFeeManagerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierProxyFeeManagerSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyFeeManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyFeeManagerSet struct {
	OldFeeManager common.Address
	NewFeeManager common.Address
	Raw           types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterFeeManagerSet(opts *bind.FilterOpts) (*VerifierProxyFeeManagerSetIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "FeeManagerSet")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyFeeManagerSetIterator{contract: _VerifierProxy.contract, event: "FeeManagerSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchFeeManagerSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyFeeManagerSet) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "FeeManagerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyFeeManagerSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "FeeManagerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VerifierProxy *VerifierProxyFilterer) ParseFeeManagerSet(log types.Log) (*VerifierProxyFeeManagerSet, error) {
	event := new(VerifierProxyFeeManagerSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "FeeManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyOwnershipTransferRequestedIterator struct {
	Event *VerifierProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierProxyOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferRequestedIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferRequested)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error) {
	event := new(VerifierProxyOwnershipTransferRequested)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyOwnershipTransferredIterator struct {
	Event *VerifierProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierProxyOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferredIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferred)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error) {
	event := new(VerifierProxyOwnershipTransferred)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyVerifierInitializedIterator struct {
	Event *VerifierProxyVerifierInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyVerifierInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyVerifierInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierProxyVerifierInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierProxyVerifierInitializedIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyVerifierInitialized struct {
	VerifierAddress common.Address
	Raw             types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterVerifierInitialized(opts *bind.FilterOpts) (*VerifierProxyVerifierInitializedIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "VerifierInitialized")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyVerifierInitializedIterator{contract: _VerifierProxy.contract, event: "VerifierInitialized", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchVerifierInitialized(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "VerifierInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyVerifierInitialized)
				if err := _VerifierProxy.contract.UnpackLog(event, "VerifierInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VerifierProxy *VerifierProxyFilterer) ParseVerifierInitialized(log types.Log) (*VerifierProxyVerifierInitialized, error) {
	event := new(VerifierProxyVerifierInitialized)
	if err := _VerifierProxy.contract.UnpackLog(event, "VerifierInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyVerifierSetIterator struct {
	Event *VerifierProxyVerifierSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyVerifierSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyVerifierSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierProxyVerifierSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierProxyVerifierSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyVerifierSet struct {
	OldConfigDigest [32]byte
	NewConfigDigest [32]byte
	VerifierAddress common.Address
	Raw             types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterVerifierSet(opts *bind.FilterOpts) (*VerifierProxyVerifierSetIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyVerifierSetIterator{contract: _VerifierProxy.contract, event: "VerifierSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchVerifierSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierSet) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyVerifierSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "VerifierSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VerifierProxy *VerifierProxyFilterer) ParseVerifierSet(log types.Log) (*VerifierProxyVerifierSet, error) {
	event := new(VerifierProxyVerifierSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "VerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyVerifierUnsetIterator struct {
	Event *VerifierProxyVerifierUnset

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyVerifierUnsetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyVerifierUnset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierProxyVerifierUnset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierProxyVerifierUnsetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyVerifierUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyVerifierUnset struct {
	ConfigDigest    [32]byte
	VerifierAddress common.Address
	Raw             types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterVerifierUnset(opts *bind.FilterOpts) (*VerifierProxyVerifierUnsetIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "VerifierUnset")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyVerifierUnsetIterator{contract: _VerifierProxy.contract, event: "VerifierUnset", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchVerifierUnset(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierUnset) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "VerifierUnset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyVerifierUnset)
				if err := _VerifierProxy.contract.UnpackLog(event, "VerifierUnset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VerifierProxy *VerifierProxyFilterer) ParseVerifierUnset(log types.Log) (*VerifierProxyVerifierUnset, error) {
	event := new(VerifierProxyVerifierUnset)
	if err := _VerifierProxy.contract.UnpackLog(event, "VerifierUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_VerifierProxy *VerifierProxy) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VerifierProxy.abi.Events["AccessControllerSet"].ID:
		return _VerifierProxy.ParseAccessControllerSet(log)
	case _VerifierProxy.abi.Events["FeeManagerSet"].ID:
		return _VerifierProxy.ParseFeeManagerSet(log)
	case _VerifierProxy.abi.Events["OwnershipTransferRequested"].ID:
		return _VerifierProxy.ParseOwnershipTransferRequested(log)
	case _VerifierProxy.abi.Events["OwnershipTransferred"].ID:
		return _VerifierProxy.ParseOwnershipTransferred(log)
	case _VerifierProxy.abi.Events["VerifierInitialized"].ID:
		return _VerifierProxy.ParseVerifierInitialized(log)
	case _VerifierProxy.abi.Events["VerifierSet"].ID:
		return _VerifierProxy.ParseVerifierSet(log)
	case _VerifierProxy.abi.Events["VerifierUnset"].ID:
		return _VerifierProxy.ParseVerifierUnset(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VerifierProxyAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x953e92b1a6442e9c3242531154a3f6f6eb00b4e9c719ba8118fa6235e4ce89b6")
}

func (VerifierProxyFeeManagerSet) Topic() common.Hash {
	return common.HexToHash("0x04628abcaa6b1674651352125cb94b65b289145bc2bc4d67720bb7d966372f03")
}

func (VerifierProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VerifierProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VerifierProxyVerifierInitialized) Topic() common.Hash {
	return common.HexToHash("0x1f2cd7c97f4d801b5efe26cc409617c1fd6c5ef786e79aacb90af40923e4e8e9")
}

func (VerifierProxyVerifierSet) Topic() common.Hash {
	return common.HexToHash("0xbeb513e532542a562ac35699e7cd9ae7d198dcd3eee15bada6c857d28ceaddcf")
}

func (VerifierProxyVerifierUnset) Topic() common.Hash {
	return common.HexToHash("0x11dc15c4b8ac2b183166cc8427e5385a5ece8308217a4217338c6a7614845c4c")
}

func (_VerifierProxy *VerifierProxy) Address() common.Address {
	return _VerifierProxy.address
}

type VerifierProxyInterface interface {
	GetVerifier(opts *bind.CallOpts, configDigest [32]byte) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SAccessController(opts *bind.CallOpts) (common.Address, error)

	SFeeManager(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	InitializeVerifier(opts *bind.TransactOpts, verifierAddress common.Address) (*types.Transaction, error)

	SetAccessController(opts *bind.TransactOpts, accessController common.Address) (*types.Transaction, error)

	SetFeeManager(opts *bind.TransactOpts, feeManager common.Address) (*types.Transaction, error)

	SetVerifier(opts *bind.TransactOpts, currentConfigDigest [32]byte, newConfigDigest [32]byte, addressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnsetVerifier(opts *bind.TransactOpts, configDigest [32]byte) (*types.Transaction, error)

	Verify(opts *bind.TransactOpts, payload []byte) (*types.Transaction, error)

	VerifyBulk(opts *bind.TransactOpts, payloads [][]byte) (*types.Transaction, error)

	FilterAccessControllerSet(opts *bind.FilterOpts) (*VerifierProxyAccessControllerSetIterator, error)

	WatchAccessControllerSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyAccessControllerSet) (event.Subscription, error)

	ParseAccessControllerSet(log types.Log) (*VerifierProxyAccessControllerSet, error)

	FilterFeeManagerSet(opts *bind.FilterOpts) (*VerifierProxyFeeManagerSetIterator, error)

	WatchFeeManagerSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyFeeManagerSet) (event.Subscription, error)

	ParseFeeManagerSet(log types.Log) (*VerifierProxyFeeManagerSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error)

	FilterVerifierInitialized(opts *bind.FilterOpts) (*VerifierProxyVerifierInitializedIterator, error)

	WatchVerifierInitialized(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierInitialized) (event.Subscription, error)

	ParseVerifierInitialized(log types.Log) (*VerifierProxyVerifierInitialized, error)

	FilterVerifierSet(opts *bind.FilterOpts) (*VerifierProxyVerifierSetIterator, error)

	WatchVerifierSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierSet) (event.Subscription, error)

	ParseVerifierSet(log types.Log) (*VerifierProxyVerifierSet, error)

	FilterVerifierUnset(opts *bind.FilterOpts) (*VerifierProxyVerifierUnsetIterator, error)

	WatchVerifierUnset(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierUnset) (event.Subscription, error)

	ParseVerifierUnset(log types.Log) (*VerifierProxyVerifierUnset, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
