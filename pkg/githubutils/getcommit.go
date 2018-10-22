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
Package githubutils provides utility functions to utilize github in the context
of retrieving and saving ethpm package manifests
*/
package githubutils

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetCommit takes a local repository directory and branch, default is 'master' if an
// empty string, and return the latest commit hash
func GetCommit(packagedir string, branch string) (commit string, err error) {
	var commithashpath string
	if packagedir == "" {
		if packagedir, err = os.Getwd(); err != nil {
			err = fmt.Errorf("Could not get working directory: '%v'", err)
			return "", err
		}
	}
	if branch == "" {
		branch = "master"
	}
	commithashpath = packagedir + "/.git/refs/heads/" + branch
	file, err := os.Open(filepath.FromSlash(commithashpath))
	if err != nil {
		err = fmt.Errorf("Could not open file '%v': '%v'", commithashpath, err)
		return "", err
	}
	info, _ := file.Stat()
	commitbytes := make([]byte, info.Size())
	_, err = file.Read(commitbytes)
	if err != nil {
		err = fmt.Errorf("Could not read file '%v': '%v'", commithashpath, err)
		return "", err
	}
	commitbytes = commitbytes[:len(commitbytes)-1]
	commit = string(commitbytes)
	return
}
