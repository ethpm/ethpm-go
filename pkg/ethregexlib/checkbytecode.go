package ethregexlib

import (
	"errors"
	"fmt"
	"regexp"
)

// CheckBytecode ensures the string is a valid hexadecimal string
func CheckBytecode(bc string) (err error) {
	re := regexp.MustCompile("^(0x|0X)[a-fA-F0-9]+$")
	matched := re.MatchString(bc)
	if !matched {
		err = errors.New("Does not conform to a hexadecimal string")
	} else if (len(bc) % 2) != 0 {
		err = fmt.Errorf("The string does not contain 2 "+
			"characters per byte, length is showing '%v'", len(bc))
	}
	return
}
