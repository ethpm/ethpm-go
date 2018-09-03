package ethcontract

import (
	"errors"
	"testing"

	bc "github.com/ethpm/ethpm-go/pkg/bytecode"
)

func TestValidate(t *testing.T) {
	var want error
	var got error
	ct := ContractType{}

	got = ct.Validate("PauloPiresToken")
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}

	ct.ContractName = " token-contract"
	got = ct.Validate("BenVeraToken")
	want = errors.New("contract_type[BenVeraToken]:contract_name error 'Name ' token-contract' " +
		"does not conform to the standard. Please check for extra whitespace and see " +
		"https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier for the requirement.'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	ct.ContractName = ""
	dub := &bc.UnlinkedBytecode{}
	dub.Bytecode = "0x3g"
	ct.DeploymentBytecode = dub
	got = ct.Validate("NickGToken")
	want = errors.New("deployment_bytecode for contract_type[NickGToken] returned the " +
		"following error: unlinked_bytecode:bytecode error 'Does not conform to a hexadecimal string'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}
	dub.Bytecode = ""

	rub := &bc.UnlinkedBytecode{}
	rub.Bytecode = "0x324"
	ct.RuntimeBytecode = rub
	got = ct.Validate("DaviToken")
	want = errors.New("runtime_bytecode for contract_type[DaviToken] returned the " +
		"following error: unlinked_bytecode:bytecode error 'The string does not contain 2 characters per byte, length is showing '5''")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	dub.Bytecode = "0x431101"
	rub.Bytecode = "0x3071d1"

	got = ct.Validate("NoMoreTokens")
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}
