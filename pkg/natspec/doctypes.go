package natspec

type DevDoc struct {
	Author       string                       `json:"author"`
	Title        string                       `json:"title"`
	Methods      map[string]map[string]string `json:"methods"`
	Invariants   []map[string]string          `json:"invariants"`
	Construction []map[string]string          `json:"construction"`
}

type UserDoc struct {
	Source          string                       `json:"source"`
	Language        string                       `json:"language"`
	LanguageVersion string                       `json:"languageVersion"`
	Methods         map[string]map[string]string `json:"methods"`
	Invariants      []map[string]string          `json:"invariants"`
	Construction    []map[string]string          `json:"construction"`
}

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
