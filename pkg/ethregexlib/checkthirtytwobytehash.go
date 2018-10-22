package ethregexlib

import (
	"errors"
	"regexp"
)

// CheckThirtyTwoByteHash ensures the string is a valid 32-byte hash
func CheckThirtyTwoByteHash(bc string) (err error) {
	re := regexp.MustCompile("^(0x|0X)[a-fA-F0-9]{64}$")
	matched := re.MatchString(bc)
	if !matched {
		err = errors.New("Does not conform to a 32-byte hash")
	}
	return
}
