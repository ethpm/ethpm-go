package librarylink

import (
	"fmt"
	"regexp"
)

// LinkReference A defined location in some bytecode which requires linking
type LinkReference struct {
	Offsets []int  `json:"offsets"`
	Length  int    `json:"length"`
	Name    string `json:"name"`
}

// CheckUniqueOffsets ensures we have unique offset values
func (l *LinkReference) CheckUniqueOffsets() (err error) {
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

// CheckName ensures name is formatted correctly
func (l *LinkReference) CheckName() (err error) {
	re := regexp.MustCompile("^[a-zA-Z][-_a-zA-Z0-9]{0,255}$")
	matched := re.MatchString(l.Name)
	if !matched {
		err = fmt.Errorf("Field 'name' does not conform to the standard. Please see " +
			"https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier " +
			"for the requirement.")
	}
	return
}
