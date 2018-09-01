package librarylink

import (
	"errors"
	"testing"
)

func TestCheckType(t *testing.T) {
	var want error
	var got error
	l := LinkValue{}

	l.Type = " literal"
	got = l.CheckType()
	want = errors.New("Field 'type' needs to be one of 'literal' or 'reference' with " +
		"no whitespace. Showing value as ' literal'")

	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	l.Type = "reference"
	got = l.CheckType()
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}

func TestCheckValue(t *testing.T) {
	var want error
	var got error
	l := LinkValue{}

	l.Type = "literal"
	l.Value = "github-diaswrd"
	got = l.CheckValue()
	want = errors.New("'type' is decalred as 'literal' and field 'value' does " +
		"not conform to a hexadecimal string")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	l.Value = "0x460t7a2e"
	got = l.CheckValue()
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	l.Value = "0x431100"
	got = l.CheckValue()
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}

	l.Type = "reference"
	got = l.CheckValue()
	want = errors.New("'type' is decalred as 'reference' and field 'value' does " +
		"not conform to the name standard defined here " +
		"https://ethpm.github.io/ethpm-spec/glossary.html#term-contract-instance")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	l.Value = "joshua-hannan-Is-the-MAN"
	got = l.CheckValue()
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}
