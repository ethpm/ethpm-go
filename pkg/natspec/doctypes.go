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
Package natspec provides DevDoc, UserDoc, and DocUnion structs which correlate
with Ethereum compiler natspec output
*/
package natspec

import "fmt"

// Method defines a method object in the doc json
type Method struct {
	Details string            `json:"details,omitempty"`
	Params  map[string]string `json:"params,omitempty"`
	Return  string            `json:"return,omitempty"`
}

// DevDoc The output of an ethereum compiler's natspec developer documentation
type DevDoc struct {
	Author       string              `json:"author,omitempty"`
	Construction []map[string]string `json:"construction,omitempty"`
	Invariants   []map[string]string `json:"invariants,omitempty"`
	Methods      map[string]*Method  `json:"methods,omitempty"`
	Title        string              `json:"title,omitempty"`
}

// UserDoc The output of am ethereum compiler's natspec user documentation
type UserDoc struct {
	Construction    []map[string]string `json:"construction,omitempty"`
	Invariants      []map[string]string `json:"invariants,omitempty"`
	Language        string              `json:"language,omitempty"`
	LanguageVersion string              `json:"languageVersion,omitempty"`
	Methods         map[string]*Method  `json:"methods,omitempty"`
	Source          string              `json:"source,omitempty"`
}

// DocUnion The union of devdoc and userdoc
type DocUnion struct {
	Author          string              `json:"author,omitempty"`
	Construction    []map[string]string `json:"construction,omitempty"`
	Invariants      []map[string]string `json:"invariants,omitempty"`
	Language        string              `json:"language,omitempty"`
	LanguageVersion string              `json:"languageVersion,omitempty"`
	Methods         map[string]*Method  `json:"methods,omitempty"`
	Source          string              `json:"source,omitempty"`
	Title           string              `json:"title,omitempty"`
}

// CreateUnion takes a DevDoc and UserDoc struct and combines them into a
// DocUnion struct
func (du *DocUnion) CreateUnion(dd *DevDoc, ud *UserDoc) {
	if ud != nil {
		du.Language = ud.Language
		du.LanguageVersion = ud.LanguageVersion
		du.Source = ud.Source
	} else {
		fmt.Println("User Docs not included in output.")
	}
	if dd != nil {
		du.Author = dd.Author
		du.Construction = dd.Construction
		du.Invariants = dd.Invariants
		du.Methods = dd.Methods
		du.Title = dd.Title
	} else {
		fmt.Println("Developer Docs not included in output.")
	}
	return
}
