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
Package bytecode provides requisite structs and utility functions to build and
validate unlinked and unlinked bytecode objects for the ethpm v2 manifest. Information
about these objects can be found here http://ethpm.github.io/ethpm-spec/package-spec.html#the-bytecode-object
*/
package bytecode

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
	liblink "github.com/ethpm/ethpm-go/pkg/librarylink"
)

// StandardJSONBC is the bytecode or deployedBytecode object from a compiler's
// standard JSON output object
type StandardJSONBC struct {
	LinkReferences map[string]map[string][]map[string]int `json:"linkReferences,omitempty"`
	Object         string                                 `json:"object,omitempty"`
}

// UnlinkedBytecode A bytecode object for unlinked bytecode.
type UnlinkedBytecode struct {
	Bytecode       string                   `json:"bytecode,omitempty"`
	LinkReferences []*liblink.LinkReference `json:"link_references,omitempty"`
}

// LinkedBytecode A bytecode object for linked bytecode
type LinkedBytecode struct {
	Bytecode         string                   `json:"bytecode,omitempty"`
	LinkDependencies []*liblink.LinkValue     `json:"link_dependencies,omitempty"`
	LinkReferences   []*liblink.LinkReference `json:"link_references,omitempty"`
}

// Build takes a compiler standard output bytecode object as a json string
// and builds the UnlinkedBytecode struct
func (ub *UnlinkedBytecode) Build(jsonstring string) (err error) {
	var s *StandardJSONBC
	var contractcount int
	if err = json.Unmarshal([]byte(jsonstring), &s); err != nil {
		err = fmt.Errorf("Error parsing standard json bytecode object: '%v'", err)
		return
	}
	if s == nil {
		err = errors.New("No unlinked bytecode received in json string")
		return
	}
	for k := range s.LinkReferences {
		contractcount += len(s.LinkReferences[k])
	}
	ub.LinkReferences = make([]*liblink.LinkReference, contractcount)
	contractcount = 0
	for k := range s.LinkReferences {
		for z, v := range s.LinkReferences[k] {
			ub.LinkReferences[contractcount] = &liblink.LinkReference{}
			ub.LinkReferences[contractcount].Build(z, v)
			s.Object = addLinkRefZeros(s.Object, ub.LinkReferences[contractcount])
			contractcount++
		}
	}
	ub.Bytecode = s.Object
	return
}

func addLinkRefZeros(bytecode string, lr *liblink.LinkReference) string {
	l := lr.Length * 2
	zeros := strings.Repeat("0", l)
	for _, x := range lr.Offsets {
		spot := x * 2
		bytecode = bytecode[:spot] + zeros + bytecode[spot+l:]
	}
	return bytecode
}

// Validate with UnlinkedBytecode ensures the UnlinkedBytecode object conforms to the standard
// described here https://ethpm.github.io/ethpm-spec/package-spec.html#bytecode
func (ub *UnlinkedBytecode) Validate() (err error) {
	if (ub.Bytecode == "") || (ub.Bytecode == "0x") {
		err = errors.New("bytecode empty and is a required field")
		return
	}
	if retErr := ethregexlib.CheckBytecode(ub.Bytecode); retErr != nil {
		err = fmt.Errorf("unlinked_bytecode:bytecode error '%v'", retErr)
		return
	}
	if retErr := checkLinkReferences(ub.Bytecode, ub.LinkReferences); retErr != nil {
		err = retErr
	}
	return
}

// Build the linked bytecode string and create a LinkedBytecode struct. Each
// dependency and reference for LinkedBytecode should be added through the
// AddLinkDependencies and AddLinkReference utility functions.
func (lb *LinkedBytecode) Build(bc string) (err error) {
	lb.Bytecode = bc
	return
}

// AddLinkDependencies will take an array of LinkValue objects and add them
// to the LinkedBytecode object. This function is not currently built into
// any compiler or deployment workflow.
func (lb *LinkedBytecode) AddLinkDependencies(lv []*liblink.LinkValue) {
	if len(lb.LinkDependencies) == 0 {
		lb.LinkDependencies = make([]*liblink.LinkValue, len(lv))
	}
	for i, v := range lv {
		lb.LinkDependencies[i] = v
	}
	return
}

// AddLinkReference will take an array of LinkReference objects and add them
// to the LinkedBytecode object. This function is not currently built into
// any compiler or deployment workflow.
func (lb *LinkedBytecode) AddLinkReference(lr []*liblink.LinkReference) {
	if len(lb.LinkReferences) == 0 {
		lb.LinkReferences = make([]*liblink.LinkReference, len(lr))
	}
	for i, v := range lr {
		lb.LinkReferences[i] = v
	}
	return
}

// Validate with LinkedBytecode ensures the LinkedBytecode object conforms to the standard
// described here https://ethpm.github.io/ethpm-spec/package-spec.html#bytecode
func (lb *LinkedBytecode) Validate(dependencyLengths map[string]int) (err error) {
	if (lb.Bytecode == "") || (lb.Bytecode == "0x") {
		err = errors.New("bytecode empty and is a required field")
		return
	}
	if retErr := ethregexlib.CheckBytecode(lb.Bytecode); retErr != nil {
		err = fmt.Errorf("linked_bytecode:bytecode error '%v'", retErr)
		return
	}
	if retErr := checkLinkReferences(lb.Bytecode, lb.LinkReferences); retErr != nil {
		err = retErr
		return
	}
	if retErr := checkLinkDependencies(lb.Bytecode, lb.LinkDependencies, dependencyLengths); retErr != nil {
		err = retErr
	}
	return
}

// checkLinkReferences validates each of the link references against the bytecode
func checkLinkReferences(bc string, lr []*liblink.LinkReference) (err error) {
	length := len(bc)
OuterLoop:
	for k, v := range lr {
		if retErr := v.Validate(); retErr != nil {
			err = fmt.Errorf("link_reference at position '%v' returned the following error: "+
				"%v+", k, retErr)
			break
		}
		for i, z := range v.Offsets {
			if (z + v.Length) >= ((length - 2) / 2) {
				err = fmt.Errorf("link_reference at position '%v' has invalid length for offset "+
					"at postion %v. Offset '%v' plus '%v' is out of bounds for the bytecode.", k, i, z, v.Length)
				break OuterLoop
			}
		}
	}
	return
}

// checkLinkDependencies validates each of the link dependencies against the link references
func checkLinkDependencies(bc string, lv []*liblink.LinkValue, depLengths map[string]int) (err error) {
	length := len(bc)
OuterLoop:
	for k, v := range lv {
		if retErr := v.Validate(depLengths); retErr != nil {
			err = fmt.Errorf("link_dependency at position '%v' returned the following error: "+
				"%v+", k, retErr)
			break
		}
		for i, z := range v.Offsets {
			if v.Type == "literal" {
				depLength := (len(v.Value) - 2) / 2
				if (z + depLength) >= ((length - 2) / 2) {
					err = fmt.Errorf("link_dependency at position '%v' has invalid length for offset "+
						"at postion %v. Offset '%v' plus '%v' (byte length of value '%v') is out of bounds "+
						"for the bytecode.", k, i, z, depLength, v.Value)
					break OuterLoop
				}
			}
			if (z + depLengths[v.Value]) >= ((length - 2) / 2) {
				err = fmt.Errorf("link_dependency at position '%v' has invalid length for offset "+
					"at postion %v. Offset '%v' plus '%v' (byte length of dependency '%v') is out of bounds "+
					"for the bytecode.", k, i, z, depLengths[v.Value], v.Value)
				break OuterLoop
			}
		}
	}
	return
}
