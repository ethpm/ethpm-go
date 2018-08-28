package manifest

// UnlinkedBytecode A bytecode object with the following key/value pairs.
type UnlinkedBytecode struct {
	Bytecode       string           `json:"bytecode,omitempty"`
	LinkReferences []*LinkReference `json:"link_references,omitempty"`
}

// LinkedBytecode A bytecode object with the following key/value pairs.
type LinkedBytecode struct {
	Bytecode         string           `json:"bytecode,omitempty"`
	LinkReferences   []*LinkReference `json:"link_references,omitempty"`
	LinkDependencies []*LinkValue     `json:"link_dependencies,omitempty"`
}
