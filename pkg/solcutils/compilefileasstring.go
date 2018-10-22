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
Package solcutils provides utilities for compiling solidity files using an
installed solc compiler.
*/
package solcutils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

// SolcOptimizer represents the optimizer settings in a standard json input
// https://solidity.readthedocs.io/en/v0.4.24/using-the-compiler.html#input-description
type SolcOptimizer struct {
	Enabled bool `json:"enabled,omitempty"`
	Runs    int  `json:"runs,omitempty"`
}

// SolcSettings represents the settings object in a standard json input
type SolcSettings struct {
	Optimizer       *SolcOptimizer                 `json:"optimizer,omitempty"`
	OutputSelection map[string]map[string][]string `json:"outputSelection,omitempty"`
	Remappings      []string                       `json:"remappings,omitempty"`
}

// StandardInput represents the standard json input object
type StandardInput struct {
	Language string                       `json:"language,omitempty"`
	Sources  map[string]map[string]string `json:"sources,omitempty"`
	Settings *SolcSettings                `json:"settings,omitempty"`
}

// CompileFileAsString takes all of the information requested and returns the
// standard input and standard output json objects as strings
func CompileFileAsString(compiler string,
	projectdir string,
	contractname string,
	inline bool,
	filepath string,
	fileasstring string,
	dependencies []string,
	optimize bool,
	runs int,
) (stdinjson string, stdoutjson string, err error) {
	so := SolcOptimizer{
		optimize,
		runs,
	}
	selections := []string{"evm.bytecode", "evm.deployedBytecode"}
	opsinner := map[string][]string{
		"*": selections,
	}
	ops := map[string]map[string][]string{
		"*": opsinner,
	}
	remappings := make([]string, len(dependencies))
	for i, v := range dependencies {
		remappings[i] = v + "/=./" + v + "/"
	}
	ss := SolcSettings{
		&so,
		ops,
		remappings,
	}
	sourcesinner := make(map[string]string)
	if inline {
		sourcesinner["content"] = fileasstring
	} else {
		sourcesinner["url"] = filepath
	}
	sources := map[string]map[string]string{
		contractname: sourcesinner,
	}
	si := StandardInput{
		"Solidity",
		sources,
		&ss,
	}

	tjson, _ := json.Marshal(si)
	tmpfile, err := ioutil.TempFile("", "tsol")
	defer os.Remove(tmpfile.Name())
	if err != nil {
		err = fmt.Errorf("Error creating temp json: '%v'", err)
		return
	}
	if _, err = tmpfile.Write(tjson); err != nil {
		err = fmt.Errorf("Error writing temp json: '%v'", err)
		return
	}
	if err = tmpfile.Close(); err != nil {
		err = fmt.Errorf("Error closing temp json: '%v'", err)
		return
	}
	execlocation, err := exec.LookPath(compiler)
	if err != nil {
		err = fmt.Errorf("Error getting solc bin location: '%v'", err)
		return
	}

	var wd string
	if projectdir == "" {
		wd, err = os.Getwd()
		if err != nil {
			err = fmt.Errorf("Error getting working directory: '%v'", err)
			return
		}
	} else {
		wd = projectdir
	}

	execCmd := exec.Command(execlocation, "--allow-paths", wd, "--standard-json")
	f, err := os.Open(tmpfile.Name())
	i, err := f.Stat()
	b := make([]byte, i.Size())
	f.Read(b)
	stdinjson = string(b)
	f.Seek(0, 0)
	if err != nil {
		err = fmt.Errorf("Error reading temp json: '%v'", err)
		return
	}
	stdin, _ := execCmd.StdinPipe()
	io.Copy(stdin, f)
	stdin.Close()
	f.Close()
	execOut, err := execCmd.Output()
	if err != nil {
		err = fmt.Errorf("Error calling exec command: '%v'", err)
		return
	}
	stdoutjson = string(execOut)
	return
}
