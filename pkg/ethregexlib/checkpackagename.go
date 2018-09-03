package ethregexlib

import (
	"errors"
	"fmt"
	"regexp"
)

// CheckPackageName ensures the string is a proper ethpm package name
func CheckPackageName(s string) (err error) {
	if s == "" {
		err = errors.New("must provide a package name")
		return
	}
	re := regexp.MustCompile("^[a-z][a-z0-9_-]{0,255}$")
	matched := re.MatchString(s)
	if !matched {
		err = fmt.Errorf("Name '%v' does not conform to the standard. Please see "+
			"https://ethpm.github.io/ethpm-spec/package-spec.html#package-name-package-name "+
			"for the spec", s)
	}
	return
}
