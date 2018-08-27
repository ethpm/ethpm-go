package manifest

type UnlinkedBytecode struct {
	Bytecode       string          `json:"bytecode"`
	LinkReferences []LinkReference `json:"linkReferences"`
}

type LinkedBytecode struct {
	Bytecode         string          `json:"bytecode"`
	LinkReferences   []LinkReference `json:"linkReferences"`
	LinkDependencies []LinkValue     `json:"linkDependencies"`
}
