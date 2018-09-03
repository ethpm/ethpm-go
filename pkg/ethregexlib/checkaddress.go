package ethregexlib

import (
	"errors"
	"regexp"
)

// CheckAddress ensures the string is a valid Ethereum address
func CheckAddress(bc string) (err error) {
	re := regexp.MustCompile("^(0x|0X)[a-fA-F0-9]{40}$")
	matched := re.MatchString(bc)
	if !matched {
		err = errors.New("Does not conform to an Ethereum address")
	}
	return
}
