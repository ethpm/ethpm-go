package bytecode

import liblink "github.com/ethpm/ethpm-go/pkg/librarylink"

// UnlinkedBytecode A bytecode object with the following key/value pairs.
type UnlinkedBytecode struct {
	Bytecode       string                   `json:"bytecode,omitempty"`
	LinkReferences []*liblink.LinkReference `json:"link_references,omitempty"`
}

// LinkedBytecode A bytecode object with the following key/value pairs.
type LinkedBytecode struct {
	Bytecode         string                   `json:"bytecode,omitempty"`
	LinkReferences   []*liblink.LinkReference `json:"link_references,omitempty"`
	LinkDependencies []*liblink.LinkValue     `json:"link_dependencies,omitempty"`
}
