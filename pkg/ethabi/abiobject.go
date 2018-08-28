package ethabi

// InputOutput An object with an input or output name and its primitive type
type InputOutput struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// ABIObject An object that conforms to the ethereum abi schema
type ABIObject struct {
	Constant        bool          `json:"constant"`
	Inputs          []InputOutput `json:"inputs"`
	Name            string        `json:"name"`
	Outputs         []InputOutput `json:"outputs"`
	Payable         bool          `json:"payable"`
	StateMutability string        `json:"stateMutability"`
	Type            string        `json:"type"`
}
