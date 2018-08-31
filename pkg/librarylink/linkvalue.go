package librarylink

// LinkValue A value for an individual link reference in a contract's bytecode
type LinkValue struct {
	Offsets []int  `json:"offsets"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}
