package bytecode

import (
	"log"
	"testing"
)

func TestSetVersion(t *testing.T) {
	c := &CompilerInformation{}
	got := c.SetVersion("solc")

	if got != nil {
		t.Fatalf("Got '%v', expected '<nil>'", got)
	}
}

func TestSetSettingsFromJSON(t *testing.T) {
	c := &CompilerInformation{}

	s := `{ "optimizer": { "enabled": true, "runs": 200 }, "outputSelection": { "*": { "*": ["abi", "evm.bytecode", "evm.deployedBytecode"] } } }`
	c.SetSettingsFromJSON(s)
	got := c.Settings.(map[string]interface{})["optimizer"].(map[string]interface{})["enabled"].(bool)

	if !got {
		t.Fatalf("Got '%v', expected 'true'", got)
	}
}

func ExampleCompilerInformation() {
	c := &CompilerInformation{}
	if err := c.SetVersion("solc"); err != nil {
		log.Fatal(err)
	}

	s := `{ "optimizer": { "enabled": true, "runs": 200 }, "outputSelection": { "*": { "*": ["abi", "evm.bytecode", "evm.deployedBytecode"] } } }`
	if err := c.SetSettingsFromJSON(s); err != nil {
		log.Fatal(err)
	}
}
