package ethpm

import (
	"errors"
	"testing"
)

func TestMetaValidate(t *testing.T) {
	var want error
	var got error

	p := PackageMeta{}
	links := make(map[string]string)
	links["documentation"] = "https://github.com/gnidan"
	links["website"] = "./github.com/Hackdom"
	p.Links = links

	got = p.Validate()
	want = errors.New("PackageMeta:links returned error 'Invalid uri contained in " +
		"key 'website'. It contains value './github.com/Hackdom', please change value to a valid absolute uri.'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	delete(links, "website")
	got = p.Validate()
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}

}
