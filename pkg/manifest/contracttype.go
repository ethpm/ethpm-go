package manifest

import "github.com/Hackdom/ethpm-go/pkg/ethabi"
import "github.com/Hackdom/ethpm-go/pkg/natspec"

type ContractType struct {
	ContractName       string           `json:"contractName"`
	DeploymentBytecode UnlinkedBytecode `json:"deploymentBytecode"`
	RuntimeBytecode    UnlinkedBytecode `json:"runtimeBytecode"`
	ABI                ethabi.ABIObject `json:"abi"`
	Natspec            natspec.DocUnion `json:"natspec"`
	Compiler           Compiler         `json:"compiler"`
}
