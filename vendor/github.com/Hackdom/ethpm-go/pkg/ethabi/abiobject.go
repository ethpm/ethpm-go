package ethabi

type InputOutput struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ABIObject struct {
	Constant        bool          `json:"constant"`
	Inputs          []InputOutput `json:"inputs"`
	Name            string        `json:"name"`
	Outputs         []InputOutput `json:"outputs"`
	Payable         bool          `json:"payable"`
	StateMutability string        `json:"stateMutability"`
	Type            string        `json:"type"`
}
