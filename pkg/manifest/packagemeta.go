package manifest

type Link struct {
	Resource string `json:"resource"`
	Uri      string `json:"uri"`
}

type PackageMeta struct {
	Authors     []string `json:"authors"`
	License     string   `json:"license"`
	Description string   `json:"description"`
	Keywords    string   `json:"keywords"`
	Links       []Link   `json:"links"`
}
