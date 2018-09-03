package librarylink

import (
	"fmt"
)

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
