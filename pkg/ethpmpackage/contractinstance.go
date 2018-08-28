package ethpmpackage

// ContractInstance Data for a deployed instance of a contract
type ContractInstance struct {
	ContractType    string               `json:"contract_type"`
	Address         string               `json:"address"`
	Transaction     string               `json:"transaction,omitempty"`
	Block           string               `json:"block,omitempty"`
	RuntimeBytecode *LinkedBytecode      `json:"runtime_bytecode,omitempty"`
	Compiler        *CompilerInformation `json:"compiler,omitempty"`
}
