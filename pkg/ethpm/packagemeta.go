package ethpm

// PackageMeta Metadata about the package
type PackageMeta struct {
	Authors     []string          `json:"authors,omitempty"`
	License     string            `json:"license,omitempty"`
	Description string            `json:"description,omitempty"`
	Keywords    string            `json:"keywords,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
}
