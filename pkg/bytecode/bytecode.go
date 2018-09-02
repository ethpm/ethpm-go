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
	LinkReferences   []*liblink.LinkReference `json:"link_references,omitempty"`
	LinkDependencies []*liblink.LinkValue     `json:"link_dependencies,omitempty"`
}

// Validate ensures the UnlinkedBytecode object conforms to the standard
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

/*
// CheckLinkedBytecode ensures the string is a valid hexadecimal string
func (lb *LinkedBytecode) CheckLinkedBytecode() (err error) {
	re := regexp.MustCompile("^(0x|0X)?[a-fA-F0-9]+$")
	matched := re.MatchString(lb.Bytecode)
	if !matched {
		err = errors.New("field 'bytecode' does not conform to a hexadecimal string")
	} else if (len(lb.Bytecode) % 2) != 0 {
		err = fmt.Errorf("field 'bytecode' is not valid, the string does not contain 2 "+
			"characters per byte, length is showing '%v'", len(lb.Bytecode))
	}
	return
}

// CheckLinkReferencesLinked validates each of the link references against the bytecode
func (lb *LinkedBytecode) CheckLinkReferencesLinked() (err error) {
	if err = lb.CheckLinkedBytecode(); err != nil {
		return
	}
	length := len(lb.Bytecode)
OuterLoop:
	for k, v := range lb.LinkReferences {
		if retErr := v.CheckName(); retErr != nil {
			err = fmt.Errorf("link_reference at position '%v' returned the following error: "+
				"%v+", k, retErr)
			break
		} else if retErr := v.CheckUniqueOffsets(); retErr != nil {
			err = fmt.Errorf("link_reference at position '%v' returned the following error: "+
				"%v+", k, retErr)
			break
		}
		for i, z := range v.Offsets {
			if z >= ((length - 2) / 2) {
				err = fmt.Errorf("link_reference at position '%v' has invalid offset at postion "+
					"%v. Value '%v' is out of bounds for the bytecode.", k, i, z)
				break OuterLoop
			} else if (z + v.Length) >= ((length - 2) / 2) {
				err = fmt.Errorf("link_reference at position '%v' has invalid length for offset "+
					"at postion %v. Offset '%v' plus '%v' is out of bounds for the bytecode.", k, i, z, v.Length)
				break OuterLoop
			}
		}
	}
	return
}

// CheckLinkDependencies validates each of the link dependencies against the link references
func (lb *LinkedBytecode) CheckLinkDependencies() (err error) {
	if err = lb.CheckLinkReferencesLinked(); err != nil {
		return
	}
	for k, v := range lb.LinkDependencies {
		if retErr := v.CheckValue(); retErr != nil {
			err = fmt.Errorf("link_dependency at position '%v' returned the following error: "+
				"%v+", k, retErr)
			break
		} else if retErr := v.CheckUniqueOffsets(); retErr != nil {
			err = fmt.Errorf("link_dependency at position '%v' returned the following error: "+
				"%v+", k, retErr)
			break
		}
		/***STOPPED HERE!!!
	}
	return
}*/
