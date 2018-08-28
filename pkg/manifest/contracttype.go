package manifest

import "github.com/Hackdom/ethpm-go/pkg/ethabi"
import "github.com/Hackdom/ethpm-go/pkg/natspec"

// ContractType Data for a contract type included in this package
type ContractType struct {
	ContractName       string               `json:"contract_name,omitempty"`
	DeploymentBytecode *UnlinkedBytecode    `json:"deployment_bytecode,omitempty"`
	RuntimeBytecode    *UnlinkedBytecode    `json:"runtime_bytecode,omitempty"`
	ABI                *ethabi.ABIObject    `json:"abi,omitempty"`
	Natspec            *natspec.DocUnion    `json:"natspec,omitempty"`
	Compiler           *CompilerInformation `json:"compiler,omitempty"`
}
