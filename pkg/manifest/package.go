package manifest

type Package struct {
	PackageName       string                                 `json:"packageName"`
	Version           string                                 `json:"version"`
	Meta              PackageMeta                            `json:"meta"`
	Sources           map[string]string                      `json:"sources"`
	ContractTypes     map[string]ContractType                `json:"contractTypes"`
	Deployments       map[string]map[string]ContractInstance `json:"deployments"`
	BuildDependencies map[string]string                      `json:"buildDependencies"`
}
