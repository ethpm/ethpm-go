package ethpm

import (
	"fmt"
)

// BuildFromManifestJSON takes a json object string and returns a PacakgeManifest struct. Minimum information
// required is package_name and version. Any additional information must conform to the
// spec.
func BuildFromManifestJSON(jsonstring string) (p PackageManifest, err error) {
	err = p.Read(jsonstring)
	if err != nil {
		err = fmt.Errorf("Could not read json string: '%v'", err)
		return
	}
	p.ManifestVersion = "2"
	err = p.Validate()
	if err != nil {
		err = fmt.Errorf("Could not build manifest: '%v'", err)
	}
	return
}

// CreateNewManifest takes a package name and version, checkes validity according
// to the ethpm v2 spec, and returns a new PackageManifest
func CreateNewManifest(packagename string, version string) (p *PackageManifest, err error) {
	p = &PackageManifest{}
	p.PackageName = packagename
	p.Version = version
	p.ManifestVersion = "2"
	if err = p.Validate(); err != nil {
		err = fmt.Errorf("Error creating manifest: '%v'", err)
	}
	return
}
