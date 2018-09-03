package ethpm

import (
	"errors"
	"testing"
)

func TestManifestValidate(t *testing.T) {
	var want error
	var got error
	p := PackageManifest{}

	p.ManifestVersion = "2"
	got = p.Validate()
	want = errors.New("PackageManifest:package_name returned error 'must provide a package name'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.PackageName = "alexs-package"
	p.ManifestVersion = " 2"
	got = p.Validate()
	want = errors.New("PackageManifest:manifest_version returned error 'manifest_version " +
		"should be 2, manifest_version is showing  2. Ensure there are no extra spaces or characters'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.ManifestVersion = "2 "
	got = p.Validate()
	want = errors.New("PackageManifest:manifest_version returned error 'manifest_version " +
		"should be 2, manifest_version is showing 2 . Ensure there are no extra spaces or characters'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.ManifestVersion = "2"
	got = p.Validate()
	want = errors.New("PackageManifest:version returned error 'must provide a version number'")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.Version = "1.0.0"
	got = p.Validate()
	if got != nil {
		t.Fatalf("Got '%v', expected '<nil>'", got)
	}
}

/*
func TestCheckSources(t *testing.T) {
	var want error
	var got error
	p := PackageManifest{}

	p.Sources = make(map[string]string)
	p.Sources["./hello"] = "/Nick"
	got = p.CheckSources()
	want = errors.New("Source with key './hello' and location value '/Nick' does not exist or is unreachable. " +
		"Please check the url or filepath and fix or consider contacting the maintainer.")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	delete(p.Sources, "./hello")
	p.Sources["/not-valid-other-chris"] = "https://github.com/modular-network"
	got = p.CheckSources()
	want = errors.New("Invalid path for source key '/not-valid-other-chris'. Please make this a relative path in accordance " +
		"with the spec found here https://ethpm.github.io/ethpm-spec/package-spec.html#sources-sources.")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	delete(p.Sources, "/not-valid-other-chris")
	p.Sources["./contracts"] = "https://github.com/modular-network/ethereum-libraries"
	got = p.CheckSources()
	if got != nil {
		t.Fatalf("Got '%v', expected '<nil>'", got)
	}
}*/
