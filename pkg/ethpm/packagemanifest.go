package ethpm

import (
	"fmt"
	"net/url"
	"os"
	"regexp"

	"github.com/ethpm/ethpm-go/pkg/ethcontract"
	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
)

// PackageManifest EthPM Manifest Specification
type PackageManifest struct {
	ManifestVersion   string                                              `json:"manifest_version"`
	PackageName       string                                              `json:"package_name"`
	Version           string                                              `json:"version"`
	Meta              *PackageMeta                                        `json:"meta,omitempty"`
	Sources           map[string]string                                   `json:"sources,omitempty"`
	ContractTypes     map[string]*ethcontract.ContractType                `json:"contract_types,omitempty"`
	Deployments       map[string]map[string]*ethcontract.ContractInstance `json:"deployments,omitempty"`
	BuildDependencies map[string]string                                   `json:"build_dependencies,omitempty"`
}

// CheckManifestVersion ensures the correct manifest version is used
func (p *PackageManifest) CheckManifestVersion() error {
	matched, err := regexp.MatchString("^2$", p.ManifestVersion)
	if (err == nil) && (!matched) {
		err = fmt.Errorf("manifest_version should be 2, manifest_version is "+
			"showing %v. Ensure there are no extra spaces or characters", p.ManifestVersion)
	}
	return err
}

// CheckPackageName ensures package_name is formatted correctly
func (p *PackageManifest) CheckPackageName() (err error) {
	re := regexp.MustCompile("^[a-z][a-z0-9_-]{0,255}$")
	matched := re.MatchString(p.PackageName)
	if !matched {
		err = fmt.Errorf("package_name does not conform to the standard. Please see " +
			"https://ethpm.github.io/ethpm-spec/package-spec.html#package-name-package-name " +
			"for the spec")
	}
	return
}

// CheckVersion ensures package versioning matches semver standard
func (p *PackageManifest) CheckVersion() (err error) {
	// go compatible verion of semver regex created by David Fichtmueller
	// here https://github.com/semver/semver/issues/232#issuecomment-405596809
	re := regexp.MustCompile("^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-(" +
		"(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-]" +
		"[0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$")
	matched := re.MatchString(p.Version)
	if !matched {
		err = fmt.Errorf("version does not conform to semver. Please check your package version string")
	}
	return
}

// CheckSources ensures the keys and values in the source mapping is formatted correctly
func (p *PackageManifest) CheckSources() (err error) {
	var uri *url.URL

	for k, v := range p.Sources {
		uri, err = url.Parse(v)
		if err != nil {
			break
		} else {
			if a := uri.IsAbs(); !a {
				_, err = os.Stat(v)
				if err != nil {
					err = fmt.Errorf("Source with key '%v' and location value '%v' does not exist or is unreachable. "+
						"Please check the url or filepath and fix or consider contacting the maintainer.", k, v)
					break
				}
			}
		}
		re := regexp.MustCompile("^(?:[\\w]\\:|\\.\\/)([a-z_\\-\\s0-9\\.]+(?:\\\\|\\/)?)+$")
		matched := re.MatchString(k)
		if !matched {
			err = fmt.Errorf("Invalid path for source key '%v'. Please make this a relative path in accordance "+
				"with the spec found here https://ethpm.github.io/ethpm-spec/package-spec.html#sources-sources.", k)
		}
	}
	return
}

func (p *PackageManifest) CheckContractTypes() (err error) {
	for k, v := range p.ContractTypes {
		if retErr := ethregexlib.CheckName(k); retErr != nil {
			err = fmt.Errorf("contract_types key '%v' does not conform to the standard. Please see "+
				"https://ethpm.github.io/ethpm-spec/glossary.html#term-contract-name "+
				"for the spec", k)
			break
		}
		if retErr := v.Validate(k); retErr != nil {
			err = fmt.Errorf("contract_type with key '%v' returned the following error: "+
				"%v+", k, retErr)
			break
		}
	}
	return
}
