// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/BLASTchain/blast/bl-bindings/solc"
)

const StorageSetterStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var StorageSetterStorageLayout = new(solc.StorageLayout)

var StorageSetterDeployedBin = "0x608060405234801561001057600080fd5b50600436106100885760003560e01c8063a6ed563e1161005b578063a6ed563e1461013a578063bd02d0f51461013a578063ca446dd914610156578063e2a4853a146100df57600080fd5b80630528afe21461008d57806321f8a721146100a25780634e91db08146100df57806354fd4d50146100f1575b600080fd5b6100a061009b3660046101d7565b610164565b005b6100b56100b036600461024c565b6101c7565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100a06100ed366004610265565b9055565b61012d6040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b6040516100d69190610287565b6101486100b036600461024c565b6040519081526020016100d6565b6100a06100ed3660046102fa565b8060005b818110156101c1576101af84848381811061018557610185610343565b905060400201600001358585848181106101a1576101a1610343565b905060400201602001359055565b806101b981610372565b915050610168565b50505050565b60006101d1825490565b92915050565b600080602083850312156101ea57600080fd5b823567ffffffffffffffff8082111561020257600080fd5b818501915085601f83011261021657600080fd5b81358181111561022557600080fd5b8660208260061b850101111561023a57600080fd5b60209290920196919550909350505050565b60006020828403121561025e57600080fd5b5035919050565b6000806040838503121561027857600080fd5b50508035926020909101359150565b600060208083528351808285015260005b818110156102b457858101830151858201604001528201610298565b818111156102c6576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b6000806040838503121561030d57600080fd5b82359150602083013573ffffffffffffffffffffffffffffffffffffffff8116811461033857600080fd5b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036103ca577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(StorageSetterStorageLayoutJSON), StorageSetterStorageLayout); err != nil {
		panic(err)
	}

	layouts["StorageSetter"] = StorageSetterStorageLayout
	deployedBytecodes["StorageSetter"] = StorageSetterDeployedBin
}
