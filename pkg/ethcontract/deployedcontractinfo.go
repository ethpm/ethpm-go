package ethcontract

import (
	liblink "github.com/ethpm/ethpm-go/pkg/librarylink"
)

// DeployedContractInfo should be built by tools during a compilation and deployment
// workflow, then send this object to the ContractInstance builder to build a
// ContractInstance for this manifest. It is not currently included in any workflow.
type DeployedContractInfo struct {
	Address      string
	Block        string
	ContractName string
	CT           *ContractType
	BC           string
	LV           []*liblink.LinkValue
	LR           []*liblink.LinkReference
	Transaction  string
}

// AddLinkValue is a helper function to add a LinkValue object to the array
func (d *DeployedContractInfo) AddLinkValue(l *liblink.LinkValue) {
	if len(d.LV) == 0 {
		d.LV = make([]*liblink.LinkValue, 1)
		d.LV[0] = l
	} else {
		d.LV = append(d.LV, l)
	}
}

// AddLinkReference is a helper function to add a LinkReference object to the array
func (d *DeployedContractInfo) AddLinkReference(l *liblink.LinkReference) {
	if len(d.LR) == 0 {
		d.LR = make([]*liblink.LinkReference, 1)
		d.LR[0] = l
	} else {
		d.LR = append(d.LR, l)
	}
}
