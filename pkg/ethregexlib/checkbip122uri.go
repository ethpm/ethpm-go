package ethregexlib

import (
	"fmt"
	"regexp"
)

// CheckBIP122URI ensures the string is formatted to the proper uri spec
// defined here https://github.com/bitcoin/bips/blob/master/bip-0122.mediawiki
func CheckBIP122URI(s string) (err error) {
	re := regexp.MustCompile("^(blockchain:)((\\/\\/[a-fA-F0-9]{64})?)\\/(tx|block|address)\\/([a-fA-F0-9]{64})$")
	matched := re.MatchString(s)
	if !matched {
		err = fmt.Errorf("String '%v' does not conform to the BIP122 URI standard. Please check for extra "+
			"whitespace and see https://github.com/bitcoin/bips/blob/master/bip-0122.mediawiki "+
			"for the requirement.", s)
	}
	return
}
