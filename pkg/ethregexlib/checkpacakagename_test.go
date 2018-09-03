package ethregexlib

import (
	"errors"
	"testing"
)

func TestCheckPackageName(t *testing.T) {
	var want error
	var got error

	s := "H"
	got = CheckPackageName(s)
	want = errors.New("Name 'H' does not conform to the standard. Please see " +
		"https://ethpm.github.io/ethpm-spec/package-spec.html#package-name-package-name " +
		"for the spec")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "2"
	got = CheckPackageName(s)
	want = errors.New("Name '2' does not conform to the standard. Please see " +
		"https://ethpm.github.io/ethpm-spec/package-spec.html#package-name-package-name " +
		"for the spec")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "hello-Me"
	got = CheckPackageName(s)
	want = errors.New("Name 'hello-Me' does not conform to the standard. Please see " +
		"https://ethpm.github.io/ethpm-spec/package-spec.html#package-name-package-name " +
		"for the spec")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "hello-piper"
	got = CheckPackageName(s)
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}
