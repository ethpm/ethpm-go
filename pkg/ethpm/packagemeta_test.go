package ethpm

import (
	"errors"
	"testing"
)

func TestCheckLinks(t *testing.T) {
	var want error
	var got error
	p := PackageManifest{}
	p.Meta = &PackageMeta{}
	links := make(map[string]string)
	links["documentation"] = "https://github.com/gnidan"
	links["website"] = "./github.com/Hackdom"
	p.Meta.Links = links

	got = p.Meta.CheckLinks()
	want = errors.New("Invalid uri contained in <meta>:<links>, key 'website' contains value './github.com/Hackdom', " +
		"please change value to a valid absolute uri.")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	delete(links, "website")
	got = p.Meta.CheckLinks()
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}

}
