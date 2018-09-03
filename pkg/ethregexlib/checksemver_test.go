package ethregexlib

import (
	"errors"
	"testing"
)

func TestCheckSemver(t *testing.T) {
	var want error
	var got error

	s := "H"
	got = CheckSemver(s)
	want = errors.New("string 'H' does not conform to semver. Please check your version string")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "0.1.0-*"
	got = CheckSemver(s)
	want = errors.New("string '0.1.0-*' does not conform to semver. Please check your version string")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = " 0"
	got = CheckSemver(s)
	want = errors.New("string ' 0' does not conform to semver. Please check your version string")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "2.0.1-commit.60cc1668"
	got = CheckSemver(s)
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}
