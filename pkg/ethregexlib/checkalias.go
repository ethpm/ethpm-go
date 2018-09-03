package ethregexlib

import (
	"fmt"
	"regexp"
)

// CheckAlias ensures the string is formatted to the spec
// defined here https://ethpm.github.io/ethpm-spec/glossary.html#term-contract-alias
func CheckAlias(s string) (err error) {
	re := regexp.MustCompile("^[a-zA-Z][-_a-zA-Z0-9]{0,255}$")
	matched := re.MatchString(s)
	if !matched {
		re = regexp.MustCompile("^[a-zA-Z][-_a-zA-Z0-9]{0,255}\\[[-a-zA-Z0-9]{1,256}\\]$")
		matched = re.MatchString(s)
		if !matched {
			err = fmt.Errorf("Name '%v' does not conform to the standard. Please check for extra "+
				"whitespace and see https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier "+
				"for the requirement.", s)
		}
	}
	return
}
