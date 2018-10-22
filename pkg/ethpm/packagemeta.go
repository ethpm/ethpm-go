package ethpm

import (
	"fmt"
	"net/url"
)

// PackageMeta Metadata about the package
type PackageMeta struct {
	Authors     []string          `json:"authors,omitempty"`
	Description string            `json:"description,omitempty"`
	Keywords    []string          `json:"keywords,omitempty"`
	License     string            `json:"license,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
}

// SetAuthors takes string arguments for package authors and sets them in the
// given PackageManifest
func (p *PackageMeta) SetAuthors(a ...string) {
	p.Authors = a
	return
}

// SetDescription takes a string argument and sets the description in the
// given PackageManifest
func (p *PackageMeta) SetDescription(d string) {
	p.Description = d
	return
}

// SetKeywords takes string arguments for keywords and sets them in the
// given PackageManifest
func (p *PackageMeta) SetKeywords(k ...string) {
	p.Keywords = k
	return
}

// SetLicense takes a string argument and sets the license in the
// given PackageManifest
func (p *PackageMeta) SetLicense(l string) {
	p.License = l
	return
}

// SetLink takes two string arguments, the first is the key that represents the
// the given link such as "website" or "documentation". The second is the uri
// for this key. This will also check validity and return an error if the uri
// is not valid
func (p *PackageMeta) SetLink(k string, uri string) (err error) {
	p.Links[k] = uri
	if err = checkLinks(p.Links); err != nil {
		err = fmt.Errorf("PackageMeta:links returned error '%v'", err)
	}
	return
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
