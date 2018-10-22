package ethregexlib

import (
	"errors"
	"testing"
)

func TestCheckBytecode(t *testing.T) {
	var want error
	var got error
	s := "34"
	got = CheckBytecode(s)
	want = errors.New("Does not conform to a hexadecimal string")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "0x3"
	got = CheckBytecode(s)
	want = errors.New("The string does not contain 2 " +
		"characters per byte, length is showing '3'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "0xh3"
	got = CheckBytecode(s)
	want = errors.New("Does not conform to a hexadecimal string")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = ""
	got = CheckBytecode(s)
	want = errors.New("Does not conform to a hexadecimal string")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	s = "0xf2"
	got = CheckBytecode(s)
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}
