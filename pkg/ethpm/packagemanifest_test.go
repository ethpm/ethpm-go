package ethpm

import (
	"errors"
	"testing"
)

func TestCheckManifestVerison(t *testing.T) {
	var want error
	var got error
	p := PackageManifest{}

	p.ManifestVersion = "2"
	got = p.CheckManifestVersion()
	if got != nil {
		t.Fatalf("Got '%v', expected nil", got)
	}

	p.ManifestVersion = " 2"
	got = p.CheckManifestVersion()
	want = errors.New("manifest_version should be 2, manifest_version is showing  2. Ensure there are no extra spaces or characters")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.ManifestVersion = "2 "
	got = p.CheckManifestVersion()
	want = errors.New("manifest_version should be 2, manifest_version is showing 2 . Ensure there are no extra spaces or characters")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}
}

func TestCheckPackageName(t *testing.T) {
	var want error
	var got error
	p := PackageManifest{}

	p.PackageName = "H"
	got = p.CheckPackageName()
	want = errors.New("package_name does not conform to the standard. Please see " +
		"https://ethpm.github.io/ethpm-spec/package-spec.html#package-name-package-name " +
		"for the spec")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.PackageName = "2"
	got = p.CheckPackageName()
	want = errors.New("package_name does not conform to the standard. Please see " +
		"https://ethpm.github.io/ethpm-spec/package-spec.html#package-name-package-name " +
		"for the spec")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.PackageName = "hello-Me"
	got = p.CheckPackageName()
	want = errors.New("package_name does not conform to the standard. Please see " +
		"https://ethpm.github.io/ethpm-spec/package-spec.html#package-name-package-name " +
		"for the spec")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.PackageName = "hello-piper"
	got = p.CheckPackageName()
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}

func TestCheckVersion(t *testing.T) {
	var want error
	var got error
	p := PackageManifest{}

	p.Version = "H"
	got = p.CheckVersion()
	want = errors.New("version does not conform to semver. Please check your package version string")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.Version = "0.1.0-*"
	got = p.CheckVersion()
	want = errors.New("version does not conform to semver. Please check your package version string")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.Version = " 0"
	got = p.CheckVersion()
	want = errors.New("version does not conform to semver. Please check your package version string")
	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	p.Version = "2.0.1-beta"
	got = p.CheckVersion()
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}

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
}
