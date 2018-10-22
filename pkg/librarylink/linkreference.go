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
Package librarylink provides the `LinkReference` and `LinkValue` structs which
describe bytecode linking locations. Further information can be found at
http://ethpm.github.io/ethpm-spec/package-spec.html#the-link-reference-object
*/
package librarylink

import (
	"fmt"

	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
)

// LinkReference A defined location in some bytecode which requires linking
type LinkReference struct {
	Offsets []int  `json:"offsets"`
	Length  int    `json:"length"`
	Name    string `json:"name,omitempty"`
}

// Build Takes a name for the linked contract, as well as the offset array from
// the compiler statndard json output link references key for this contract, and
// builds the LinkReference struct
func (l *LinkReference) Build(name string, offsets []map[string]int) {
	l.Name = name
	if len(offsets) > 0 {
		l.Length = offsets[0]["length"]
		l.Offsets = make([]int, len(offsets))
		for i, v := range offsets {
			l.Offsets[i] = v["start"]
		}
	}
	return
}

// Validate ensures the LinkReference struct conforms to the standard found
// here https://ethpm.github.io/ethpm-spec/package-spec.html#the-link-reference-object
func (l *LinkReference) Validate() (err error) {
	if retErr := checkUniqueOffsets(l.Offsets, l.Length); retErr != nil {
		err = retErr
		return
	}
	if l.Name != "" {
		if retErr := ethregexlib.CheckContractName(l.Name); retErr != nil {
			err = fmt.Errorf("LinkReference:name error '%v'", retErr)
		}
	}
	return
}
