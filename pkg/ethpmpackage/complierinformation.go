package ethpmpackage

// CompilerInformation Information about the software that was used to compile a contract type or instance
type CompilerInformation struct {
	Name     string      `json:"name"`
	Version  string      `json:"version"`
	Settings interface{} `json:"settings,omitempty"`
}
