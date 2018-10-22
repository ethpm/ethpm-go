package librarylink

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
)

// LinkValue A value for an individual link reference in a contract's bytecode
type LinkValue struct {
	Offsets []int  `json:"offsets"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

// Build Takes a name for the linked contract, as well as the offset array for
// this contract, and builds the LinkValue struct
func (l *LinkValue) Build(typeofvalue string, value string, offsets []int) {
	l.Type = typeofvalue
	l.Value = value
	l.Offsets = offsets
	return
}

// Validate ensures the LinkValue struct conforms to the standard found
// here https://ethpm.github.io/ethpm-spec/package-spec.html#the-link-value-object
func (l *LinkValue) Validate(depLengths map[string]int) (err error) {
	if len(l.Type) == 0 {
		err = errors.New("link_value does not contain 'type' and is required")
		return
	}
	if retErr := checkType(l.Type); retErr != nil {
		err = retErr
		return
	}
	if len(l.Value) == 0 {
		err = errors.New("link_value does not contain 'value' and is required")
		return
	}
	if retErr := checkValue(l.Type, l.Value); retErr != nil {
		err = retErr
		return
	}
	if len(l.Offsets) == 0 {
		err = errors.New("link_value does not contain any offsets and is required")
		return
	}
	if l.Type == "literal" {
		if retErr := checkUniqueOffsets(l.Offsets, (len(l.Value)-2)/2); retErr != nil {
			err = retErr
			return
		}
		if retErr := checkUniqueOffsets(l.Offsets, depLengths[l.Value]); retErr != nil {
			err = retErr
		}
	}
	return
}

// checkType ensures type is a proper value
func checkType(s string) (err error) {
	re := regexp.MustCompile("^(literal|reference)$")
	matched := re.MatchString(s)
	if !matched {
		err = fmt.Errorf("Field 'type' needs to be one of 'literal' or 'reference' with "+
			"no whitespace. Showing value as '%v'", s)
	}
	return
}

// checkValue ensures value complies with the defined type
func checkValue(t string, s string) (err error) {
	if t == "literal" {
		if retErr := ethregexlib.CheckBytecode(s); retErr != nil {
			err = fmt.Errorf("LinkValue:value of type 'literal' error: '%v'", retErr)
		}
	} else {
		if retErr := ethregexlib.CheckDependencyTree(s); retErr != nil {
			err = fmt.Errorf("LinkValue:value of type 'reference' error: '%v'", retErr)
		}
	}
	return
}
