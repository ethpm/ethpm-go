package ethregexlib

import (
	"errors"
	"fmt"
	"regexp"
)

// CheckSemver ensures package versioning matches semver standard
func CheckSemver(s string) (err error) {
	if s == "" {
		err = errors.New("must provide a version number")
		return
	}
	// go compatible verion of semver regex created by David Fichtmueller
	// here https://github.com/semver/semver/issues/232#issuecomment-405596809
	re := regexp.MustCompile("^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-(" +
		"(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-]" +
		"[0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$")
	matched := re.MatchString(s)
	if !matched {
		err = fmt.Errorf("string '%v' does not conform to semver. Please check your version string", s)
	}
	return
}
