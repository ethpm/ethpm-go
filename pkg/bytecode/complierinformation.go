package bytecode

import (
	"errors"
	"fmt"

	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
)

// CompilerInformation Information about the software that was used to compile
// a contract type or instance
type CompilerInformation struct {
	Name     string      `json:"name"`
	Settings interface{} `json:"settings,omitempty"`
	Version  string      `json:"version"`
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
