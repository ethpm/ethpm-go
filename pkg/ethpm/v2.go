package ethpm

import "encoding/json"

const version = "2"

func (p *PackageManifest) Read(s string) (err error) {
	jsonBytes := []byte(s)
	err = json.Unmarshal(jsonBytes, p)
	return
}

func (p *PackageManifest) Write() (s string, err error) {
	p.ManifestVersion = version

	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return
	}
	s = string(jsonBytes)
	return
}
