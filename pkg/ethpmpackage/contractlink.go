package ethpmpackage

// LinkReference A defined location in some bytecode which requires linking
type LinkReference struct {
	Offsets []int  `json:"offsets"`
	Length  int    `json:"length"`
	Name    string `json:"name"`
}

// LinkValue A value for an individual link reference in a contract's bytecode
type LinkValue struct {
	Offsets []int  `json:"offsets"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}
