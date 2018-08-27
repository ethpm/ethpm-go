package manifest

type Compiler struct {
	Name     string      `json:"name"`
	Version  string      `json:"version"`
	Settings interface{} `json:"settings"`
}
