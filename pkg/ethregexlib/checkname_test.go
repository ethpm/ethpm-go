package ethregexlib

import (
	"errors"
	"testing"
)

func TestCheckContractName(t *testing.T) {
	var want error
	var got error
	s := " H"
	got = CheckContractName(s)
	want = errors.New("Name ' H' does not conform to the standard. Please check for extra " +
		"whitespace and see https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier " +
		"for the requirement.")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "2"
	got = CheckContractName(s)
	want = errors.New("Name '2' does not conform to the standard. Please check for extra " +
		"whitespace and see https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier " +
		"for the requirement.")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "hello-piper"
	got = CheckContractName(s)
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}
