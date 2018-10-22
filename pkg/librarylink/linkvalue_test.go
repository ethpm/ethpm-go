package librarylink

import (
	"errors"
	"testing"
)

func TestLVValidate(t *testing.T) {
	var want error
	var got error
	l := LinkValue{}
	l.Offsets = []int{0}
	var deps map[string]int
	got = l.Validate(deps)
	want = errors.New("link_value does not contain 'type' and is required")

	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	l.Type = " reference"
	got = l.Validate(deps)
	want = errors.New("Field 'type' needs to be one of 'literal' or 'reference' with " +
		"no whitespace. Showing value as ' reference'")

	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	l.Type = "literal"
	l.Value = "x9"
	got = l.Validate(deps)
	want = errors.New("LinkValue:value of type 'literal' error: 'Does not conform to a hexadecimal string'")

	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	l.Value = "0x2f"
	got = l.Validate(deps)
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}
