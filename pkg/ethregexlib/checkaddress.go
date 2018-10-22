/*
The MIT License (MIT)
https://github.com/ethpm/ethpm-go/blob/master/LICENSE

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/*
Package ethregexlib provides various regex utility functions that are relevant
to ethpm and Ethereum in general. Information about this spec can be found here
http://ethpm.github.io/ethpm-spec/package-spec.html
*/
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
