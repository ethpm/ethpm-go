package gethutils

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

// GetAccountByAddress takes an address object and the geth data directory, then
// returns the key store and an account object.
func GetAccountByAddress(addr common.Address, datadir string) (ks *keystore.KeyStore, a accounts.Account, err error) {
	ks = keystore.NewKeyStore(datadir+"/keystore", 262144, 1)
	as := ks.Accounts()
	for _, v := range as {
		if v.Address == addr {
			a = v
			break
		}
	}
	if len(addr) == 0 {
		err = fmt.Errorf("No key with address '%v'", addr)
	}
	return
}
