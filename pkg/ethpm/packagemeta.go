package ethpm

import (
	"fmt"
	"net/url"
)

// PackageMeta Metadata about the package
type PackageMeta struct {
	Authors     []string          `json:"authors,omitempty"`
	Description string            `json:"description,omitempty"`
	Keywords    string            `json:"keywords,omitempty"`
	License     string            `json:"license,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
}

// Validate ensures PackageManifest conforms to the standard defined here
// https://ethpm.github.io/ethpm-spec/package-spec.html#the-package-meta-object
func (p *PackageMeta) Validate() (err error) {
	if retErr := checkLinks(p.Links); retErr != nil {
		err = fmt.Errorf("PackageMeta:links returned error '%v'", retErr)
	}
	return
}

// CheckLinks determines if valid uris are contained in Links mapping
func checkLinks(l map[string]string) (err error) {
	var uri *url.URL

	for k, v := range l {
		uri, err = url.Parse(v)
		if err != nil {
			break
		} else {
			if a := uri.IsAbs(); !a {
				err = fmt.Errorf("Invalid uri contained in key '%v'. It contains value '%v', "+
					"please change value to a valid absolute uri.", k, v)
				break
			}
		}
	}
	return
}
