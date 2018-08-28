package manifest

// PackageManifest EthPM Manifest Specification
type PackageManifest struct {
	ManifestVersion   string                                  `json:"manifest_version"`
	PackageName       string                                  `json:"package_name"`
	Version           string                                  `json:"version"`
	Meta              *PackageMeta                            `json:"meta,omitempty"`
	Sources           map[string]string                       `json:"sources,omitempty"`
	ContractTypes     map[string]*ContractType                `json:"contract_types,omitempty"`
	Deployments       map[string]map[string]*ContractInstance `json:"deployments,omitempty"`
	BuildDependencies map[string]string                       `json:"build_dependencies,omitempty"`
}
