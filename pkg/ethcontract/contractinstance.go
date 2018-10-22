package ethcontract

import (
	"fmt"
	"regexp"

	bc "github.com/ethpm/ethpm-go/pkg/bytecode"
	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
)

// ContractInstance Data for a deployed instance of a contract
type ContractInstance struct {
	Address         string                  `json:"address"`
	Block           string                  `json:"block,omitempty"`
	Compiler        *bc.CompilerInformation `json:"compiler,omitempty"`
	ContractType    string                  `json:"contract_type"`
	RuntimeBytecode *bc.LinkedBytecode      `json:"runtime_bytecode,omitempty"`
	Transaction     string                  `json:"transaction,omitempty"`
}

// Build takes a DeployedContractInfo object and creates a ContractInstance
// object. This function is not currently built into any compiler or deployment
// workflow.
func (ci *ContractInstance) Build(i *DeployedContractInfo) {
	ci.Address = i.Address
	ci.Block = i.Block
	ci.ContractType = i.ContractName
	ci.Transaction = i.Transaction
	ci.Compiler = i.CT.Compiler
	ci.RuntimeBytecode = &bc.LinkedBytecode{}
	ci.RuntimeBytecode.Build(i.BC)
	ci.RuntimeBytecode.AddLinkReference(i.LR)
	ci.RuntimeBytecode.AddLinkDependencies(i.LV)
}

// Validate ensures ContractInstance conforms to the standard defined here
// https://ethpm.github.io/ethpm-spec/package-spec.html#the-contract-instance-object
func (ci *ContractInstance) Validate(name string, dependencyLengths map[string]int) (err error) {
	if ci.ContractType == "" {
		err = fmt.Errorf("ContractInstance[%v]:contract_type is required and showing empty string", name)
		return
	}
	if retErr := checkContractType(ci.ContractType); retErr != nil {
		err = fmt.Errorf("ContractInstance[%v]:contract_type returned error '%v'", name, retErr)
		return
	}
	if retErr := ethregexlib.CheckAddress(ci.Address); retErr != nil {
		err = fmt.Errorf("ContractInstance[%v]:address error '%v'", name, retErr)
		return
	}
	if ci.Transaction != "" {
		if retErr := ethregexlib.CheckThirtyTwoByteHash(ci.Transaction); retErr != nil {
			err = fmt.Errorf("ContractInstance[%v]:transaction error '%v'", name, retErr)
			return
		}
	}
	if ci.Block != "" {
		if retErr := ethregexlib.CheckThirtyTwoByteHash(ci.Block); retErr != nil {
			err = fmt.Errorf("ContractInstance[%v]:block error '%v'", name, retErr)
			return
		}
	}
	if (ci.RuntimeBytecode != nil) && (ci.RuntimeBytecode.Bytecode != "") {
		if retErr := ci.RuntimeBytecode.Validate(dependencyLengths); retErr != nil {
			err = fmt.Errorf("ContractInstance[%v]:runtime_bytecode error '%v'", name, retErr)
			return
		}
	}
	if (ci.Compiler != nil) && (ci.Compiler.Name != "") {
		if retErr := ci.Compiler.Validate(); retErr != nil {
			err = fmt.Errorf("ContractInstance[%v]:compiler object error '%v'", name, retErr)
			return
		}
	}
	return
}

func checkContractType(s string) (err error) {
	re := regexp.MustCompile("^(.{0,256}):?[a-zA-Z][-_a-zA-Z0-9]{0,255}$")
	matched := re.MatchString(s)
	if !matched {
		re = regexp.MustCompile("^(.{0,256}):?a-zA-Z][-_a-zA-Z0-9]{0,255}\\[[-a-zA-Z0-9]{1,256}\\]$")
		matched = re.MatchString(s)
		if !matched {
			err = fmt.Errorf("contract_type '%v' does not conform to the standard. Please check for extra "+
				"whitespace and see https://ethpm.github.io/ethpm-spec/package-spec.html#contract-type-contract-type "+
				"for the requirement.", s)
		}
	}
	return
}
