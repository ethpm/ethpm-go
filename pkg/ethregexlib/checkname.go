package ethregexlib

import (
	"fmt"
	"regexp"
)

// CheckName ensures names and identifiers are formatted correctly to the spec
// defined here https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier
func CheckName(s string) (err error) {
	re := regexp.MustCompile("^[a-zA-Z][-_a-zA-Z0-9]{0,255}$")
	matched := re.MatchString(s)
	if !matched {
		err = fmt.Errorf("Does not conform to the standard. Please see " +
			"https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier " +
			"for the requirement.")
	}
	return
}
