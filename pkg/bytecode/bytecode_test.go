package bytecode

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	var want error
	var got error
	ub := UnlinkedBytecode{}

	got = ub.Validate()
	want = errors.New("bytecode empty and is a required field")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	ub.Bytecode = "0x4g"
	got = ub.Validate()
	want = errors.New("unlinked_bytecode:bytecode error 'Does not conform to a hexadecimal string'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	ub.Bytecode = "0x"
	got = ub.Validate()
	want = errors.New("bytecode empty and is a required field")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	ub.Bytecode = "0xf3"
	got = ub.Validate()
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}
