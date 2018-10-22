package ethpm

import (
	"fmt"
	"log"
	"testing"
)

func TestBuildFromManifestJSON(t *testing.T) {
	manifestjson := `{"package_name":"testpackage", "version":"0.0.1"}`
	got, err := BuildFromManifestJSON(manifestjson)
	if err != nil {
		t.Fatal(err)
	}

	want := &PackageManifest{
		ManifestVersion: "2",
		PackageName:     "testpackage",
		Version:         "0.0.1",
	}
	if got.PackageName != want.PackageName {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}
}

func ExampleBuildFromManifestJSON() {
	manifestjson := `{"package_name":"mypackage", "version":"0.0.1"}`
	packagemanifeststruct, err := BuildFromManifestJSON(manifestjson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(packagemanifeststruct.PackageName)
	// Output: mypackage
}

func TestCreateNewManifest(t *testing.T) {
	got, err := CreateNewManifest("testpackage", "4.2.0")
	if err != nil {
		t.Fatal(err)
	}

	want := &PackageManifest{
		ManifestVersion: "2",
		PackageName:     "testpackage",
		Version:         "4.2.0",
	}
	if got.PackageName != want.PackageName {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}
}

func ExampleCreateNewManifest() {
	packagemanifeststruct, err := CreateNewManifest("mypackage", "4.2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(packagemanifeststruct.PackageName)
	// Output: mypackage
}
