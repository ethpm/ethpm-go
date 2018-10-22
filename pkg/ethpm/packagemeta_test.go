package ethpm

import (
	"errors"
	"fmt"
	"testing"
)

func TestSetAuthors(t *testing.T) {
	p := PackageMeta{}
	p.SetAuthors("test", "me")
	got := p.Authors[0]
	want := "test"
	if got != want {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}
}

func ExamplePackageMeta_SetAuthors() {
	p := PackageMeta{}
	p.SetAuthors("Joshua", "Hannan")
	fmt.Println(p.Authors[0])
	// Output: Joshua
}

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
