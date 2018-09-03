package ethregexlib

import (
	"fmt"
	"regexp"
)

// CheckDependencyTree ensures the string is in proper ethpm format specified
// here https://ethpm.github.io/ethpm-spec/package-spec.html#value-value
func CheckDependencyTree(s string) (err error) {
	re := regexp.MustCompile("^(([a-zA-Z][-a-zA-Z0-9_]{0,255})|(([a-z][a-z0-9_-]{0,255}):)+([a-zA-Z][-a-zA-Z0-9_]{0,255}))$")
	matched := re.MatchString(s)
	if !matched {
		err = fmt.Errorf("Name '%v' does not conform to the dependency tree standard. Please "+
			"check for whitespace and see https://ethpm.github.io/ethpm-spec/package-spec.html#value-value "+
			"for the spec", s)
	}
	return
}
