package natspec

// DevDoc The output of an ethereum compiler's natspec developer documentation
type DevDoc struct {
	Author       string                       `json:"author"`
	Title        string                       `json:"title"`
	Methods      map[string]map[string]string `json:"methods"`
	Invariants   []map[string]string          `json:"invariants"`
	Construction []map[string]string          `json:"construction"`
}

// UserDoc The output of am ethereum compiler's natspec user documentation
type UserDoc struct {
	Source          string                       `json:"source"`
	Language        string                       `json:"language"`
	LanguageVersion string                       `json:"languageVersion"`
	Methods         map[string]map[string]string `json:"methods"`
	Invariants      []map[string]string          `json:"invariants"`
	Construction    []map[string]string          `json:"construction"`
}

// DocUnion The union of devdoc and userdoc
type DocUnion struct {
	Author          string                       `json:"author"`
	Title           string                       `json:"title"`
	Methods         map[string]map[string]string `json:"methods"`
	Source          string                       `json:"source"`
	Language        string                       `json:"language"`
	LanguageVersion string                       `json:"languageVersion"`
	Invariants      []map[string]string          `json:"invariants"`
	Construction    []map[string]string          `json:"construction"`
}
