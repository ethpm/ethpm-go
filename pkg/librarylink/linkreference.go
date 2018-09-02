package librarylink

import (
	"fmt"

	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
)

// LinkReference A defined location in some bytecode which requires linking
type LinkReference struct {
	Offsets []int  `json:"offsets"`
	Length  int    `json:"length"`
	Name    string `json:"name,omitempty"`
}

// Validate ensures the LinkReference struct conforms to the standard found
// here https://ethpm.github.io/ethpm-spec/package-spec.html#the-link-reference-object
func (l *LinkReference) Validate() (err error) {
	if retErr := checkUniqueOffsets(l.Offsets, l.Length); retErr != nil {
		err = retErr
		return
	}
	if l.Name != "" {
		if retErr := ethregexlib.CheckName(l.Name); retErr != nil {
			err = fmt.Errorf("LinkReference:name error '%v'", retErr)
		}
	}
	return
}

// checkUniqueOffsets ensures offset values plus length do not overlap
func checkUniqueOffsets(o []int, l int) (err error) {
	encountered := make(map[int]bool)
OuterLoop:
	for k, v := range o {
		for e := range encountered {
			if (v >= e) && (v <= (e + l)) {
				err = fmt.Errorf("'Offsets' must contain unique values, '%v' appears at "+
					"index key '%v' and is within a byte range already used, '%v'-'%v'", v, k, e, (e + l))
				break OuterLoop
			}
		}
		encountered[v] = true
	}
	return
}
