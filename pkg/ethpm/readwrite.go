package ethpm

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

// WriteToDisk takes a PackageManifest struct, validates, and writes it to the
// location defined by directoryname. If directoryname is an empty string, it
// writes to the current working directory.
func (p *PackageManifest) WriteToDisk(directoryname string) (err error) {
	var pm *os.File

	if err = p.Validate(); err != nil {
		err = fmt.Errorf("PackageManifest not properly formatted: '%v'", err)
		return
	}

	properjson, _ := p.Write()

	if directoryname == "" {
		if directoryname, err = os.Getwd(); err != nil {
			err = fmt.Errorf("Could not get working directory: '%v'", err)
			return
		}
	}
	f := filepath.Join(directoryname, "ethpm.json")

	if pm, err = os.Open(f); os.IsNotExist(err) {
		err = nil
		pm, err = os.Create(directoryname + "ethpm.json")
		if err != nil {
			err = fmt.Errorf("Could not create file ethpm.json: '%v'", err)
			return
		}
	} else if err != nil {
		err = fmt.Errorf("Could not open existing ethpm.json: '%v'", err)
		return
	}
	defer pm.Close()

	_, err = pm.WriteString(properjson)
	if err != nil {
		err = fmt.Errorf("Could not write ethpm.json: '%v'", err)
		return
	}
	pm.Sync()
	return
}
