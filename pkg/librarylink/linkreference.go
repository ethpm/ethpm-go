package librarylink

// LinkReference A defined location in some bytecode which requires linking
type LinkReference struct {
	Offsets []int  `json:"offsets"`
	Length  int    `json:"length"`
	Name    string `json:"name"`
}
