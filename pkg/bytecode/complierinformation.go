package bytecode

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
)

// CompilerInformation Information about the software that was used to compile
// a contract type or instance
type CompilerInformation struct {
	Name     string      `json:"name"`
	Settings interface{} `json:"settings,omitempty"`
	Version  string      `json:"version"`
}

// Build takes the compiler name, such as "solc", and the 'settings' object from the
// standard input json and builds the CompilerInformation object
func (c *CompilerInformation) Build(compiler string, settingsjsonstring string) (err error) {
	if err = c.SetVersion(compiler); err != nil {
		err = fmt.Errorf("Error building compiler information: '%v'", err)
	} else if err = c.SetSettingsFromJSON(settingsjsonstring); err != nil {
		err = fmt.Errorf("Error building compiler information: '%v'", err)
	} else {
		c.Name = compiler
	}
	return
}

// SetName sets the name of the compiler, such as "solc"
func (c *CompilerInformation) SetName(t string) {
	c.Name = t
	return
}

// SetSettingsFromJSON accepts the 'settings' key from solc standard input json
// as a string and assigns it to the PackageManifest compiler information
func (c *CompilerInformation) SetSettingsFromJSON(jsonstring string) (err error) {
	var i map[string]interface{}

	jsonBytes := []byte(jsonstring)
	if err = json.Unmarshal(jsonBytes, &i); err != nil {
		err = fmt.Errorf("Error getting settings from JSON: '%v'", err)
	}
	c.Settings = i
	return
}

// SetVersion takes the name of the compiler and sets the version. The compiler
// must be installed on your system. Currently only tested against "solc --version"
// output
func (c *CompilerInformation) SetVersion(t string) (err error) {
	c.Version, err = GetVersion(t)
	return
}

// GetVersion takes the name of the compiler and returns the version being used.
// Currently only tested against "solc --version" output
func GetVersion(t string) (version string, err error) {
	execlocation, err := exec.LookPath(t)
	if err != nil {
		err = fmt.Errorf("Error getting solc bin location: '%v'", err)
		return
	}

	execCmd := exec.Command(execlocation, "--version")

	execOut, err := execCmd.Output()
	if err != nil {
		err = fmt.Errorf("Error calling exec command: '%v'", err)
		return
	}

	versionarray := strings.Split(strings.Split(string(execOut), "Version: ")[1], ".")
	version = strings.Trim(strings.Replace(fmt.Sprint(versionarray[:len(versionarray)-2]), " ", ".", -1), "[]")
	return
}

// Validate ensures ContractType conforms to the standard defined here
// https://ethpm.github.io/ethpm-spec/package-spec.html#the-compiler-information-object
func (c *CompilerInformation) Validate() (err error) {
	if c.Name == "" {
		err = errors.New("CompilerInformation:name is required and showing empty string")
		return
	}
	if retErr := ethregexlib.CheckSemver(c.Version); retErr != nil {
		err = fmt.Errorf("CompilerInformation:version returned the following error '%v'", retErr)
	}
	return
}
