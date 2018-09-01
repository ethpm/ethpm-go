package librarylink

import (
	"errors"
	"fmt"
	"regexp"
)

// LinkValue A value for an individual link reference in a contract's bytecode
type LinkValue struct {
	Offsets []int  `json:"offsets"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

// CheckUniqueOffsets ensures we have unique offset values
func (l *LinkValue) CheckUniqueOffsets() (err error) {
	encountered := map[int]int{}

	for k, v := range l.Offsets {
		if encountered[v] != 0 {
			err = fmt.Errorf("'Offsets' must contain unique values, '%v' appears at "+
				"index keys '%v' and '%v'", v, encountered[v]-1, k)
			break
		}
		encountered[v] = k + 1
	}
	return
}

// CheckType ensures type is a proper value
func (l *LinkValue) CheckType() (err error) {
	re := regexp.MustCompile("^(literal|reference)$")
	matched := re.MatchString(l.Type)
	if !matched {
		err = fmt.Errorf("Field 'type' needs to be one of 'literal' or 'reference' with "+
			"no whitespace. Showing value as '%v'", l.Type)
	}
	return
}

// CheckValue ensures value complies with the defined type
func (l *LinkValue) CheckValue() (err error) {
	if err = l.CheckType(); err != nil {
		return
	}
	if l.Type == "literal" {
		re := regexp.MustCompile("^(0x|0X)[a-fA-F0-9]+$")
		matched := re.MatchString(l.Value)
		if !matched {
			err = errors.New("'type' is decalred as 'literal' and field 'value' does " +
				"not conform to a hexadecimal string")
		}
	} else {
		re := regexp.MustCompile("^[a-zA-Z][-a-zA-Z0-9_]{0,255}$")
		matched := re.MatchString(l.Value)
		if !matched {
			err = fmt.Errorf("'type' is decalred as 'reference' and field 'value' does " +
				"not conform to the name standard defined here " +
				"https://ethpm.github.io/ethpm-spec/glossary.html#term-contract-instance")
		}
	}
	return
}
