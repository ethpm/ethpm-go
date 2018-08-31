package ethcontract

import (
	bc "github.com/ethpm/ethpm-go/pkg/bytecode"
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
