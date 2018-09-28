package bytecode

import (
	"errors"
	"fmt"

	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
	liblink "github.com/ethpm/ethpm-go/pkg/librarylink"
)

// UnlinkedBytecode A bytecode object with the following key/value pairs.
type UnlinkedBytecode struct {
	Bytecode       string                   `json:"bytecode,omitempty"`
	LinkReferences []*liblink.LinkReference `json:"link_references,omitempty"`
}

// LinkedBytecode A bytecode object with the following key/value pairs.
type LinkedBytecode struct {
	Bytecode         string                   `json:"bytecode,omitempty"`
	LinkDependencies []*liblink.LinkValue     `json:"link_dependencies,omitempty"`
	LinkReferences   []*liblink.LinkReference `json:"link_references,omitempty"`
}

// Validate with UnlinkedBytecode ensures the UnlinkedBytecode object conforms to the standard
// described here https://ethpm.github.io/ethpm-spec/package-spec.html#bytecode
func (ub *UnlinkedBytecode) Validate() (err error) {
	if (ub.Bytecode == "") || (ub.Bytecode == "0x") {
		err = errors.New("bytecode empty and is a required field")
		return
	}
	if retErr := ethregexlib.CheckBytecode(ub.Bytecode); retErr != nil {
		err = fmt.Errorf("unlinked_bytecode:bytecode error '%v'", retErr)
		return
	}
	if retErr := checkLinkReferences(ub.Bytecode, ub.LinkReferences); retErr != nil {
		err = retErr
	}
	return
}

// Validate with LinkedBytecode ensures the LinkedBytecode object conforms to the standard
// described here https://ethpm.github.io/ethpm-spec/package-spec.html#bytecode
func (lb *LinkedBytecode) Validate(dependencyLengths map[string]int) (err error) {
	if (lb.Bytecode == "") || (lb.Bytecode == "0x") {
		err = errors.New("bytecode empty and is a required field")
		return
	}
	if retErr := ethregexlib.CheckBytecode(lb.Bytecode); retErr != nil {
		err = fmt.Errorf("linked_bytecode:bytecode error '%v'", retErr)
		return
	}
	if retErr := checkLinkReferences(lb.Bytecode, lb.LinkReferences); retErr != nil {
		err = retErr
		return
	}
	if retErr := checkLinkDependencies(lb.Bytecode, lb.LinkDependencies, dependencyLengths); retErr != nil {
		err = retErr
	}
	return
}

// checkLinkReferences validates each of the link references against the bytecode
func checkLinkReferences(bc string, lr []*liblink.LinkReference) (err error) {
	length := len(bc)
OuterLoop:
	for k, v := range lr {
		if retErr := v.Validate(); retErr != nil {
			err = fmt.Errorf("link_reference at position '%v' returned the following error: "+
				"%v+", k, retErr)
			break
		}
		for i, z := range v.Offsets {
			if (z + v.Length) >= ((length - 2) / 2) {
				err = fmt.Errorf("link_reference at position '%v' has invalid length for offset "+
					"at postion %v. Offset '%v' plus '%v' is out of bounds for the bytecode.", k, i, z, v.Length)
				break OuterLoop
			}
		}
	}
	return
}

// checkLinkDependencies validates each of the link dependencies against the link references
func checkLinkDependencies(bc string, lv []*liblink.LinkValue, depLengths map[string]int) (err error) {
	length := len(bc)
OuterLoop:
	for k, v := range lv {
		if retErr := v.Validate(depLengths); retErr != nil {
			err = fmt.Errorf("link_dependency at position '%v' returned the following error: "+
				"%v+", k, retErr)
			break
		}
		for i, z := range v.Offsets {
			if v.Type == "literal" {
				depLength := (len(v.Value) - 2) / 2
				if (z + depLength) >= ((length - 2) / 2) {
					err = fmt.Errorf("link_dependency at position '%v' has invalid length for offset "+
						"at postion %v. Offset '%v' plus '%v' (byte length of value '%v') is out of bounds "+
						"for the bytecode.", k, i, z, depLength, v.Value)
					break OuterLoop
				}
			}
			if (z + depLengths[v.Value]) >= ((length - 2) / 2) {
				err = fmt.Errorf("link_dependency at position '%v' has invalid length for offset "+
					"at postion %v. Offset '%v' plus '%v' (byte length of dependency '%v') is out of bounds "+
					"for the bytecode.", k, i, z, depLengths[v.Value], v.Value)
				break OuterLoop
			}
		}
	}
	return
}
