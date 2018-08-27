package manifest

type ContractInstance struct {
  ContractType    string          `json:"contractType"`
  Address         string          `json:"address"`
  Transaction     string          `json:"transaction"`
  Block           string          `json:"block"`
  RuntimeBytecode LinkedBytecode  `json:"runtimeBytecode"`
  Compiler        Compiler        `json:"compiler"`
}
