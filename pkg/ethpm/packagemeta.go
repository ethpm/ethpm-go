package ethpm

import (
	"fmt"
	"net/url"
)

// PackageMeta Metadata about the package
type PackageMeta struct {
	Authors     []string          `json:"authors,omitempty"`
	License     string            `json:"license,omitempty"`
	Description string            `json:"description,omitempty"`
	Keywords    string            `json:"keywords,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
}

// CheckLinks determines if valid uris are contained in Links mapping
func (pm *PackageMeta) CheckLinks() (err error) {
	var uri *url.URL

	for k, v := range pm.Links {
		uri, err = url.Parse(v)
		if err != nil {
			break
		} else {
			if a := uri.IsAbs(); !a {
				err = fmt.Errorf("Invalid uri contained in <meta>:<links>, key '%v' contains value '%v', "+
					"please change value to a valid absolute uri.", k, v)
				break
			}
		}
	}
	return
}
