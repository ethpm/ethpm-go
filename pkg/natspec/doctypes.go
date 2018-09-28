package natspec

// DevDoc The output of an ethereum compiler's natspec developer documentation
type DevDoc struct {
	Author       string                       `json:"author"`
	Construction []map[string]string          `json:"construction"`
	Invariants   []map[string]string          `json:"invariants"`
	Methods      map[string]map[string]string `json:"methods"`
	Title        string                       `json:"title"`
}

// UserDoc The output of am ethereum compiler's natspec user documentation
type UserDoc struct {
	Construction    []map[string]string          `json:"construction"`
	Invariants      []map[string]string          `json:"invariants"`
	Language        string                       `json:"language"`
	LanguageVersion string                       `json:"languageVersion"`
	Methods         map[string]map[string]string `json:"methods"`
	Source          string                       `json:"source"`
}

// DocUnion The union of devdoc and userdoc
type DocUnion struct {
	Author          string                       `json:"author"`
	Construction    []map[string]string          `json:"construction"`
	Invariants      []map[string]string          `json:"invariants"`
	Language        string                       `json:"language"`
	LanguageVersion string                       `json:"languageVersion"`
	Methods         map[string]map[string]string `json:"methods"`
	Source          string                       `json:"source"`
	Title           string                       `json:"title"`
}
