package ethpm

import (
	"github.com/ethpm/ethpm-go/pkg/ethcontract"
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
