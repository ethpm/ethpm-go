package ethcontract

import (
	bc "github.com/ethpm/ethpm-go/pkg/bytecode"
)

// ContractInstance Data for a deployed instance of a contract
type ContractInstance struct {
	ContractType    string                  `json:"contract_type"`
	Address         string                  `json:"address"`
	Transaction     string                  `json:"transaction,omitempty"`
	Block           string                  `json:"block,omitempty"`
	RuntimeBytecode *bc.LinkedBytecode      `json:"runtime_bytecode,omitempty"`
	Compiler        *bc.CompilerInformation `json:"compiler,omitempty"`
}
