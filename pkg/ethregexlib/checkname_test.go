package ethregexlib

import (
	"errors"
	"testing"
)

func TestCheckName(t *testing.T) {
	var want error
	var got error
	s := " H"
	got = CheckName(s)
	want = errors.New("Does not conform to the standard. Please see " +
		"https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier " +
		"for the requirement.")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "2"
	got = CheckName(s)
	want = errors.New("Does not conform to the standard. Please see " +
		"https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier " +
		"for the requirement.")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "hello-piper"
	got = CheckName(s)
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}
