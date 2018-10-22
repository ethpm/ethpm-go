/*
The MIT License (MIT)
https://github.com/ethpm/ethpm-go/blob/master/LICENSE

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/*
Package gethutils provides utility functions that use ethereum-go. Further information
for ethereum-go can be found at https://github.com/ethereum/go-ethereum
*/
package gethutils

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

// ConnectGeth takes the full path to the geth data directory in use and connects
// via the ipc connection. Geth must be active and connected to a network.
func ConnectGeth(datadir string) (*ethclient.Client, *big.Int, error) {
	var ec *ethclient.Client
	var t *big.Int
	var err error
	ec, err = ethclient.Dial(datadir + "/geth.ipc")
	if err != nil {
		err = fmt.Errorf("Could not find geth connection to '%v': '%v'", datadir+"/geth.ipc", err)
		return nil, nil, err
	}
	t, err = ec.NetworkID(context.Background())
	if err != nil {
		err = fmt.Errorf("Geth ain't acting right: '%v'", err)
		return nil, nil, err
	}
	return ec, t, nil
}
