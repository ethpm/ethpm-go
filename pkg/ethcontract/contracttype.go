package ethcontract

import (
	"fmt"

	bc "github.com/ethpm/ethpm-go/pkg/bytecode"
	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
	"github.com/ethpm/ethpm-go/pkg/natspec"
)

// ContractType Data for a contract type included in this package
type ContractType struct {
	ContractName       string                  `json:"contract_name,omitempty"`
	DeploymentBytecode *bc.UnlinkedBytecode    `json:"deployment_bytecode,omitempty"`
	RuntimeBytecode    *bc.UnlinkedBytecode    `json:"runtime_bytecode,omitempty"`
	ABI                *ABIObject              `json:"abi,omitempty"`
	Natspec            *natspec.DocUnion       `json:"natspec,omitempty"`
	Compiler           *bc.CompilerInformation `json:"compiler,omitempty"`
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
