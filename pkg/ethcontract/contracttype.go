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
	if retErr := ethregexlib.CheckName(ct.ContractName); retErr != nil {
		err = fmt.Errorf("contract_type:contract_name error '%v'", retErr)
		return
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
	if (dbc.Bytecode == "") || (dbc.Bytecode == "0x") {
		fmt.Println("No deployment_bytecode for " + name)
		return
	}
	if retErr := dbc.Validate(); retErr != nil {
		err = fmt.Errorf("deployment_bytecode returned the following error: "+
			"%v+", retErr)
	}
	return
}

func checkRuntimeBytecode(name string, rbc *bc.UnlinkedBytecode) (err error) {
	if (rbc.Bytecode == "") || (rbc.Bytecode == "0x") {
		fmt.Println("No runtime_bytecode for " + name)
		return
	}
	if retErr := rbc.Validate(); retErr != nil {
		err = fmt.Errorf("runtime_bytecode returned the following error: "+
			"%v+", retErr)
	}
	return
}
