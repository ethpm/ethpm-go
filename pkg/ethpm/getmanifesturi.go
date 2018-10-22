package ethpm

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethpm/ethpm-go/pkg/gethutils"
	"github.com/ethpm/ethpm-go/pkg/packageregistry"
)

// GetManifestURI uses an ipc connection to a locally running geth node
func GetManifestURI(repositoryaddressashex string, packagename string, version string, chainname string, gethdatadir string) (uri string, err error) {
	var nh common.Hash
	var vh common.Hash
	var ch common.Hash
	var info struct {
		Name        string
		Version     string
		ManifestURI string
	}
	hw := sha3.NewKeccak256()

	hw.Write([]byte(packagename))
	hw.Sum(nh[:0])
	hw.Reset()

	hw.Write([]byte(version))
	hw.Sum(vh[:0])
	hw.Reset()

	fh := append(nh[:], vh[:]...)
	hw.Write([]byte(fh))
	hw.Sum(ch[:0])

	br := bytes.NewReader(packageregistry.GetPackageRegistryABI())
	ethabi, _ := abi.JSON(br)
	nb, _ := ethabi.Pack("getReleaseData", ch)

	ra := common.HexToAddress(repositoryaddressashex)
	m := ethereum.CallMsg{
		From:     common.HexToAddress("0x0"),
		To:       &ra,
		Gas:      0,
		Value:    big.NewInt(0),
		GasPrice: big.NewInt(0),
		Data:     nb,
	}
	if gethdatadir == "" {
		gethdatadir = node.DefaultDataDir()
	}
	if (chainname != "") && (chainname != "mainnet") {
		gethdatadir += "/" + chainname
	}
	ec, _, err := gethutils.ConnectGeth(gethdatadir)
	if err != nil {
		fmt.Println(err)
		return
	}
	if b, e := ec.CallContract(context.Background(), m, nil); e != nil {
		fmt.Println(e)
	} else {
		if te := ethabi.Unpack(&info, "getReleaseData", b); te != nil {
			fmt.Println(te)
		}
		fmt.Println(info)
	}

	return
}
