package manifest

type LinkReference struct {
	Offsets []int  `json:"offsets"`
	Length  int    `json:"length"`
	Name    string `json:"name"`
}

type LinkValue struct {
	Offsets []int  `json:"offsets"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}
