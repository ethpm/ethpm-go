/*
The MIT License (MIT)
https://github.com/ethpm/ethpm-go/blob/master/LICENSE

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/*
Package ethcontract provides `ABIObject`, which correlates with a compiler's abi
output, as well as `ContractInstance` and `ContractType` which follows the EthPM
v2 spec for these objects. Information about these objects can be found here
http://ethpm.github.io/ethpm-spec/package-spec.html#the-contract-type-object
*/
package ethcontract

import (
	"encoding/json"
	"fmt"

	bc "github.com/ethpm/ethpm-go/pkg/bytecode"
	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
	"github.com/ethpm/ethpm-go/pkg/natspec"
)

// ContractType Data for a contract type included in this package
type ContractType struct {
	ABI                []*ABIObject            `json:"abi,omitempty"`
	Compiler           *bc.CompilerInformation `json:"compiler,omitempty"`
	ContractName       string                  `json:"contract_name,omitempty"`
	DeploymentBytecode *bc.UnlinkedBytecode    `json:"deployment_bytecode,omitempty"`
	Natspec            *natspec.DocUnion       `json:"natspec,omitempty"`
	RuntimeBytecode    *bc.UnlinkedBytecode    `json:"runtime_bytecode,omitempty"`
}

// Build takes the name of the compiler used (currently only tested with solc),
// the settings object from a standard json input, https://solidity.readthedocs.io/en/v0.4.24/using-the-compiler.html#input-description,
// as a string, and the compiler standard json ouput as a string. It then builds
// a contract type object
func (ct *ContractType) Build(compiler string, settingsjsonstring string, compileroutputjson string) (err error) {
	var i map[string]interface{}
	var b map[string]interface{}
	var dd *natspec.DevDoc
	var ud *natspec.UserDoc

	jsonBytes := []byte(compileroutputjson)
	if err = json.Unmarshal(jsonBytes, &i); err != nil {
		err = fmt.Errorf("Error getting contract type from JSON: '%v'", err)
		return
	}
	abibytes, _ := json.Marshal(i["abi"])
	devdocbytes, _ := json.Marshal(i["devdoc"])
	userdocbytes, _ := json.Marshal(i["userdoc"])
	evmbytes, _ := json.Marshal(i["evm"])
	if err = json.Unmarshal(evmbytes, &b); err != nil {
		err = fmt.Errorf("Error parsing bytes from evm object: '%v'", err)
		return
	}
	depbytecodebytes, _ := json.Marshal(b["bytecode"])
	runbytecodebytes, _ := json.Marshal(b["deployedBytecode"])

	if err = json.Unmarshal(abibytes, &ct.ABI); err != nil {
		err = fmt.Errorf("Error generating ABI: '%v'", err)
		return
	}

	if err = json.Unmarshal(devdocbytes, &dd); err != nil {
		err = fmt.Errorf("Error generating DevDoc: '%v'", err)
		return
	}

	if err = json.Unmarshal(userdocbytes, &ud); err != nil {
		err = fmt.Errorf("Error generating UserDoc: '%v'", err)
		return
	}
	ct.Natspec = &natspec.DocUnion{}
	ct.Compiler = &bc.CompilerInformation{}
	ct.DeploymentBytecode = &bc.UnlinkedBytecode{}
	ct.RuntimeBytecode = &bc.UnlinkedBytecode{}

	ct.Natspec.CreateUnion(dd, ud)
	ct.Compiler.Build(compiler, settingsjsonstring)
	ct.DeploymentBytecode.Build(string(depbytecodebytes))
	ct.RuntimeBytecode.Build(string(runbytecodebytes))

	return
}

// Validate ensures ContractType conforms to the standard defined here
// https://ethpm.github.io/ethpm-spec/package-spec.html#contract-type-object
func (ct *ContractType) Validate(name string) (err error) {
	if ct.ContractName != "" {
		if retErr := ethregexlib.CheckContractName(ct.ContractName); retErr != nil {
			err = fmt.Errorf("contract_type[%v]:contract_name error '%v'", name, retErr)
			return
		}
	}
	if retErr := checkDeploymentBytecode(name, ct.DeploymentBytecode); retErr != nil {
		err = retErr
		return
	}
	if retErr := checkRuntimeBytecode(name, ct.RuntimeBytecode); retErr != nil {
		err = retErr
	}
	return
}

// checkDeploymentBytecode ensures a proper UnlinkedBytecode object is in the ContractType struct
func checkDeploymentBytecode(name string, dbc *bc.UnlinkedBytecode) (err error) {
	if dbc != nil {
		if (dbc.Bytecode == "") || (dbc.Bytecode == "0x") {
			fmt.Printf("No deployment_bytecode for contract_type[%v]", name)
			return
		}
		if retErr := dbc.Validate(); retErr != nil {
			err = fmt.Errorf("deployment_bytecode for contract_type[%v] returned the following error: "+
				"%v", name, retErr)
		}
	}
	return
}

func checkRuntimeBytecode(name string, rbc *bc.UnlinkedBytecode) (err error) {
	if rbc != nil {
		if (rbc.Bytecode == "") || (rbc.Bytecode == "0x") {
			fmt.Printf("No runtime_bytecode for contract_type[%v]", name)
			return
		}
		if retErr := rbc.Validate(); retErr != nil {
			err = fmt.Errorf("runtime_bytecode for contract_type[%v] returned the following error: "+
				"%v", name, retErr)
		}
	}
	return
}
