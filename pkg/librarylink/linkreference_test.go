package librarylink

import (
	"errors"
	"testing"
)

func TestCheckName(t *testing.T) {
	var want error
	var got error
	l := LinkReference{}

	l.Name = "he!!0-Will"
	got = l.CheckName()
	want = errors.New("Field 'name' does not conform to the standard. Please see " +
		"https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier " +
		"for the requirement.")

	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	l.Name = "hello-Me"
	got = l.CheckName()
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}
