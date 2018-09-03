package ethpm

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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

// Validate ensures PackageManifest conforms to the standard defined here
// https://ethpm.github.io/ethpm-spec/package-spec.html#document-specification
func (p *PackageManifest) Validate() (err error) {
	if retErr := checkManifestVersion(p.ManifestVersion); retErr != nil {
		err = fmt.Errorf("PackageManifest:manifest_version returned error '%v'", retErr)
		return
	}
	if retErr := ethregexlib.CheckPackageName(p.PackageName); retErr != nil {
		err = fmt.Errorf("PackageManifest:package_name returned error '%v'", retErr)
		return
	}
	if retErr := ethregexlib.CheckSemver(p.Version); retErr != nil {
		err = fmt.Errorf("PackageManifest:version returned error '%v'", retErr)
		return
	}
	if p.Meta != nil {
		if retErr := p.Meta.Validate(); retErr != nil {
			err = fmt.Errorf("PackageManifest:meta returned error '%v'", retErr)
			return
		}
	}
	if retErr := checkSources(p.Sources); retErr != nil {
		err = fmt.Errorf("PackageManifest:sources returned error '%v'", retErr)
		return
	}
	if retErr := checkContractTypes(p.ContractTypes); retErr != nil {
		err = fmt.Errorf("PackageManifest:contract_types returned error '%v'", retErr)
		return
	}
	if retErr := checkDeployments(p.Deployments); retErr != nil {
		err = fmt.Errorf("PackageManifest:deployments returned error '%v'", retErr)
		return
	}
	if retErr := checkBuildDependencies(p.BuildDependencies); retErr != nil {
		err = fmt.Errorf("PackageManifest:build_dependencies returned error '%v'", retErr)
	}
	return
}

// checkManifestVersion ensures the correct manifest version is used
func checkManifestVersion(s string) error {
	matched, err := regexp.MatchString("^2$", s)
	if (err == nil) && (!matched) {
		err = fmt.Errorf("manifest_version should be 2, manifest_version is "+
			"showing %v. Ensure there are no extra spaces or characters", s)
	}
	return err
}

// checkSources ensures the keys and values in the source mapping is formatted correctly
func checkSources(s map[string]string) (err error) {
	var uri *url.URL

	for k, v := range s {
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

func checkContractTypes(ct map[string]*ethcontract.ContractType) (err error) {
	for k, v := range ct {
		if retErr := ethregexlib.CheckAlias(k); retErr != nil {
			err = fmt.Errorf("contract_types key '%v' does not conform to the standard. Please see "+
				"https://ethpm.github.io/ethpm-spec/glossary.html#term-contract-alias "+
				"for the spec", k)
			break
		}
		if retErr := v.Validate(k); retErr != nil {
			err = fmt.Errorf("contract_type with key '%v' returned the following error: "+
				"%v", k, retErr)
			break
		}
	}
	return
}

func checkDeployments(d map[string]map[string]*ethcontract.ContractInstance) (err error) {
OuterLoop:
	for k, v := range d {
		if retErr := ethregexlib.CheckBIP122URI(k); retErr != nil {
			err = fmt.Errorf("deployment with key '%v' returned the following error: "+
				"%v", k, retErr)
			break
		}
		for i, z := range v {
			if retErr := ethregexlib.CheckContractName(i); retErr != nil {
				err = fmt.Errorf("deployment[%v] with key '%v' returned the following error: "+
					"%v", k, i, retErr)
				break OuterLoop
			}
			dependencyLengths := make(map[string]int)
			for _, y := range z.RuntimeBytecode.LinkDependencies {
				if y.Type == "reference" {
					dependencyLengths[y.Value], err = getLinkValueDependencyLength(k, v, y.Value)
					if err != nil {
						err = fmt.Errorf("deployment[%v]:contract_instance[%v] returned the following dependency "+
							"error for dependency '%v'. Please ensure you have the dependency installed correctly: "+
							"%v", k, i, y.Value, err)
						break OuterLoop
					}
				}
			}
			if retErr := z.Validate(i, dependencyLengths); retErr != nil {
				err = fmt.Errorf("deployment[%v]:contract_instance[%v] returned the following error: "+
					"%v", k, i, retErr)
				break OuterLoop
			}
		}
	}
	return
}

func getLinkValueDependencyLength(blockchainURI string, thisDeps map[string]*ethcontract.ContractInstance, d string) (length int, err error) {
	if d == "" {
		return
	}

	var currentDir string

	if retErr := ethregexlib.CheckDependencyTree(d); retErr != nil {
		err = retErr
		return
	}
	if currentDir, err = os.Getwd(); err != nil {
		return
	}
	depTree := strings.Split(d, ":")
	treeLength := len(depTree)
	var depPath strings.Builder
	depPath.WriteString(currentDir)
	if treeLength > 1 {
		for i := 0; i < (treeLength - 2); i++ {
			depPath.WriteString("/ethpm-dependencies/" + depTree[i])
		}
		depPath.WriteString("/" + depTree[treeLength-2] + ".json")

		var file *os.File
		if file, err = os.Open(filepath.FromSlash(depPath.String())); err != nil {
			return
		}
		manifest := make([]byte, 100)
		_, err = file.Read(manifest)
		if err != nil {
			return
		}
		pm := PackageManifest{}
		err = json.Unmarshal(manifest, &pm)
		if err != nil {
			return
		}
		bc := pm.Deployments[blockchainURI][depTree[treeLength-1]].Address
		return ((len(bc) - 2) / 2), nil
	}
	bc := thisDeps[depTree[0]].Address
	return ((len(bc) - 2) / 2), nil
}

func checkBuildDependencies(bd map[string]string) (err error) {
	var uri *url.URL

	for k, v := range bd {
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
				"with the spec found here https://ethpm.github.io/ethpm-spec/package-spec.html#build-dependencies-build-dependencies.", k)
		}
	}
	return
}
