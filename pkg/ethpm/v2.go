package ethpm

import (
	"encoding/json"
)

const version = "2"

// Read will read a json string representing the package manifest
func (p *PackageManifest) Read(s string) (err error) {
	jsonBytes := []byte(s)
	err = json.Unmarshal(jsonBytes, p)
	return
}

// Write will convert the PackageManifest struct into a json string
func (p *PackageManifest) Write() (s string, err error) {
	p.ManifestVersion = version

	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return
	}
	s = string(jsonBytes)
	return
}
