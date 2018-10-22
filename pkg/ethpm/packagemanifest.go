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
Package ethpm provides the primary manifest object defined in `packagemanifest.go`.
`manifestinterface.go` defines a basic interface for a manifest object. We define
the v2 instance which implements this interface in `packagemanifest.go`.
Information about this spec can be found here
http://ethpm.github.io/ethpm-spec/package-spec.html
*/
package ethpm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethpm/ethpm-go/pkg/ethcontract"
	"github.com/ethpm/ethpm-go/pkg/ethregexlib"
	"github.com/ethpm/ethpm-go/pkg/gethutils"
	"github.com/ethpm/ethpm-go/pkg/packageregistry"
	"github.com/ethpm/ethpm-go/pkg/solcutils"
)

// PackageManifest EthPM Manifest Specification
type PackageManifest struct {
	BuildDependencies map[string]string                                   `json:"build_dependencies,omitempty"`
	ContractTypes     map[string]*ethcontract.ContractType                `json:"contract_types,omitempty"`
	Deployments       map[string]map[string]*ethcontract.ContractInstance `json:"deployments,omitempty"`
	ManifestVersion   string                                              `json:"manifest_version"`
	Meta              *PackageMeta                                        `json:"meta,omitempty"`
	PackageName       string                                              `json:"package_name"`
	Sources           map[string]string                                   `json:"sources,omitempty"`
	Version           string                                              `json:"version"`
}

// AddDependency takes the name of another package and its uri, then adds it
// to this manifest's BuildDependencies
func (p *PackageManifest) AddDependency(name string, uri string) {
	if len(p.BuildDependencies) == 0 {
		p.BuildDependencies = make(map[string]string)
	}
	p.BuildDependencies[name] = uri
	return
}

// AddContractType takes the name of the compiler installed on your system and being used,
// the settings object from the standard JSON input, https://solidity.readthedocs.io/en/v0.4.24/using-the-compiler.html#input-description,
// the standard JSON output, and the contract name. It then adds the contract type
// to this manifest.
func (p *PackageManifest) AddContractType(compiler string, settingsjsonstring string, compileroutputjson string, contractname string) (err error) {
	var i map[string]map[string]map[string]interface{}

	if len(p.ContractTypes) == 0 {
		p.ContractTypes = make(map[string]*ethcontract.ContractType)
	}

	jsonBytes := []byte(compileroutputjson)
	if err = json.Unmarshal(jsonBytes, &i); err != nil {
		err = fmt.Errorf("Error getting contract type from JSON for '%v': '%v'", contractname, err)
		return
	}
	for _, v := range i["contracts"] {
		if v[contractname] != nil {
			contractbytes, _ := json.Marshal(v[contractname])
			p.ContractTypes[contractname] = &ethcontract.ContractType{}
			p.ContractTypes[contractname].Build(compiler, settingsjsonstring, string(contractbytes))
		}
	}
	return
}

// AddDeployment takes a blockchain uri for a deployed contract instance, a
// DeployedContractInfo object, and creates a new deployment object for this
// package. This function is not currently implemented in any workflow.
func (p *PackageManifest) AddDeployment(blockchainuri string, d *ethcontract.DeployedContractInfo) {
	if len(p.Deployments) == 0 {
		p.Deployments = make(map[string]map[string]*ethcontract.ContractInstance)
	}
	p.Deployments[blockchainuri][d.ContractName].Build(d)
	return
}

// SourceInliner takes the directory containing contract files and the file type
// such as "sol", then adds this source to the package manifest. The source relative
// path should contain the path relative to the manifest json location. It can
// be an empty string, in which case, the path will be the same directory.
func (p *PackageManifest) SourceInliner(contractdir string, sourcerelativepath string, sourcetype string) (err error) {
	if contractdir == "" {
		if contractdir, err = os.Getwd(); err != nil {
			err = fmt.Errorf("Could not get working directory: '%v'", err)
			return
		}
	}
	if sourcerelativepath == "" {
		sourcerelativepath = "./"
	}
	if len(p.Sources) == 0 {
		p.Sources = make(map[string]string)
	}
	files, err := ioutil.ReadDir(contractdir)
	if err != nil {
		err = fmt.Errorf("Could not get read directory: '%v'", err)
		return
	}
	for _, f := range files {
		n := f.Name()
		if n[len(n)-3:] == sourcetype {
			if b, thiserr := ioutil.ReadFile(filepath.Join(contractdir, n)); thiserr != nil {
				err = fmt.Errorf("Could not get read %v: '%v'", n, thiserr)
			} else {
				p.Sources[sourcerelativepath+n] = string(b)
			}
		}
	}
	return
}

// AddLocalPathForSource takes the contract directory, which can be left empty and
// the current working directory will be used, the source path relative to the location
// of the manifest file, which can be empty and will be the same directory, and the
// the source type, generally .sol. It will then add a source file location as relative
// to the manifest.
func (p *PackageManifest) AddLocalPathForSource(contractdir string, sourcerelativepath string, sourcetype string) (err error) {
	if contractdir == "" {
		if contractdir, err = os.Getwd(); err != nil {
			err = fmt.Errorf("Could not get working directory: '%v'", err)
			return
		}
	}
	if sourcerelativepath == "" {
		sourcerelativepath = "./"
	}
	if len(p.Sources) == 0 {
		p.Sources = make(map[string]string)
	}
	files, err := ioutil.ReadDir(contractdir)
	if err != nil {
		err = fmt.Errorf("Could not get read directory: '%v'", err)
		return
	}
	for _, f := range files {
		n := f.Name()
		if n[len(n)-3:] == sourcetype {
			p.Sources[sourcerelativepath+n] = sourcerelativepath + n
		}
	}
	return
}

// CompileAndValidateSource takes the name of the installed compiler, such as
// solc, the project directory, contract name, if the source is inlined in the
// package manifest, set inline to true, the source file path, if source is inline
// this should be equal to the key identifying the source, the full file path to source
// if not inline, compiler optimize setting (true or false), and the number of runs
// for the optimizer (will be ignored if optimize is fale). It will then compile
// the provided contract and compare to the equivalent contract type in the manifest.
// If it is a match, then valid will return true, if not, it should return false.
// producedobject is the string representation of the generated contract type.
//
// This has not been incorporated into any workflow nor rigorously tested.
func (p *PackageManifest) CompileAndValidateSource(compiler string,
	projectdir string,
	contractname string,
	inline bool,
	filepath string,
	optimize bool,
	runs int,
) (valid bool, producedobject string, err error) {
	var fileasstring string
	if inline {
		if len(p.Sources[filepath]) == 0 {
			err = fmt.Errorf("Invalid inline source key: '%v'", err)
			return
		}
		fileasstring = p.Sources[filepath]
	}
	dependencies := make([]string, 0)
	for k := range p.BuildDependencies {
		dependencies = append(dependencies, k)
	}
	stdinjson, stdoutjson, err := solcutils.CompileFileAsString(compiler,
		projectdir,
		contractname,
		inline,
		filepath,
		fileasstring,
		dependencies,
		optimize,
		runs)
	if err != nil {
		err = fmt.Errorf("Error compiling source: '%v'", err)
		return
	}

	var s map[string]interface{}
	jsonBytes := []byte(stdinjson)
	if err = json.Unmarshal(jsonBytes, &s); err != nil {
		err = fmt.Errorf("Error getting setting from standard JSON input: '%v'", err)
		return
	}
	settingsbytes, _ := json.Marshal(s["settings"])
	ec := &ethcontract.ContractType{}
	err = ec.Build(compiler, string(settingsbytes), stdoutjson)
	if err != nil {
		err = fmt.Errorf("Error building the contracty type object: '%v'", err)
		return
	}
	b, _ := json.Marshal(ec)
	producedobject = string(b)
	if ec.DeploymentBytecode.Bytecode == p.ContractTypes[contractname].DeploymentBytecode.Bytecode {
		valid = true
	}
	return
}

// PublishToRepositoryWithPassword uses an ipc connection with a locally running
// geth node. It takes an onchain repository address for the connected network,
// the manifest's uri, the wallet address you wish to use in the local keystore,
// the preferred gas price, chain name (ie rinkeby), and the geth data directory
// if other than default, if the default is used, it can be an empty string. It will
// then publish this package in the repository referred to.
//
// This function has not been incorporated into any workflow
func (p *PackageManifest) PublishToRepositoryWithPassword(repositoryaddressashex string,
	manifesturi string,
	fromaddressashex string,
	gaspriceinwei int64,
	chainname string,
	gethdatadir string,
) (err error) {
	var bgas *big.Int

	ra := common.HexToAddress(repositoryaddressashex)
	fa := common.HexToAddress(fromaddressashex)
	if gethdatadir == "" {
		gethdatadir = node.DefaultDataDir()
	}
	if (chainname != "") && (chainname != "mainnet") {
		gethdatadir += "/" + chainname
	}
	ec, networkID, err := gethutils.ConnectGeth(gethdatadir)
	if err != nil {
		err = fmt.Errorf("Error connecting to geth: '%v'", err)
		return
	}
	ks, a, err := gethutils.GetAccountByAddress(fa, gethdatadir)
	if err != nil {
		err = fmt.Errorf("Error getting wallet: '%v'", err)
		return
	}
	nonce, err := ec.NonceAt(context.Background(), fa, nil)
	if err != nil {
		err = fmt.Errorf("Error getting wallet nonce: '%v'", err)
		return
	}
	if gaspriceinwei == 0 {
		bgas, err = ec.SuggestGasPrice(context.Background())
		if err != nil {
			err = fmt.Errorf("Error getting suggested gas: '%v'", err)
			return
		}
	} else {
		bgas = big.NewInt(gaspriceinwei)
	}
	br := bytes.NewReader(packageregistry.GetPackageRegistryABI())
	ethabi, _ := abi.JSON(br)
	nb, _ := ethabi.Pack("release", p.PackageName, p.Version, manifesturi)

	tx := types.NewTransaction(nonce, ra, big.NewInt(0), 100000, bgas, nb)

	password := gethutils.GetPassword()
	stx, err := ks.SignTxWithPassphrase(a, password, tx, networkID)

	if err != nil {
		err = fmt.Errorf("Signing failed: '%v'", err)
		return
	}
	err = ec.SendTransaction(context.Background(), stx)
	if err != nil {
		err = fmt.Errorf("Internal error: '%v'", err)
	}
	return
}

// Validate ensures PackageManifest conforms to the standard defined here
// https://ethpm.github.io/ethpm-spec/package-spec.html#document-specification
func (p *PackageManifest) Validate() (err error) {
	if retErr := checkManifestVersion(p.ManifestVersion); retErr != nil {
		err = fmt.Errorf("PackageManifest:manifest_version returned error '%v'", retErr)
		return
	}
	if retErr := ethregexlib.CheckPackageName(p.PackageName); retErr != nil {
		err = fmt.Errorf("PackageManifest:package_name returned error '%v'", retErr)
		return
	}
	if retErr := ethregexlib.CheckSemver(p.Version); retErr != nil {
		err = fmt.Errorf("PackageManifest:version returned error '%v'", retErr)
		return
	}
	if p.Meta != nil {
		if retErr := p.Meta.Validate(); retErr != nil {
			err = fmt.Errorf("PackageManifest:meta returned error '%v'", retErr)
			return
		}
	}
	if retErr := checkSources(p.Sources); retErr != nil {
		err = fmt.Errorf("PackageManifest:sources returned error '%v'", retErr)
		return
	}
	if retErr := checkContractTypes(p.ContractTypes); retErr != nil {
		err = fmt.Errorf("PackageManifest:contract_types returned error '%v'", retErr)
		return
	}
	if retErr := checkDeployments(p.Deployments); retErr != nil {
		err = fmt.Errorf("PackageManifest:deployments returned error '%v'", retErr)
		return
	}
	if retErr := checkBuildDependencies(p.BuildDependencies); retErr != nil {
		err = fmt.Errorf("PackageManifest:build_dependencies returned error '%v'", retErr)
	}
	return
}

// checkManifestVersion ensures the correct manifest version is used
func checkManifestVersion(s string) error {
	matched, err := regexp.MatchString("^2$", s)
	if (err == nil) && (!matched) {
		err = fmt.Errorf("manifest_version should be 2, manifest_version is "+
			"showing %v. Ensure there are no extra spaces or characters", s)
	}
	return err
}

// checkSources ensures the keys and values in the source mapping is formatted correctly
func checkSources(s map[string]string) (err error) {
	var uri *url.URL

	for k, v := range s {
		uri, err = url.Parse(v)
		if err != nil {
			break
		} else {
			if a := uri.IsAbs(); !a {
				_, err = os.Stat(v)
				if err != nil {
					err = fmt.Errorf("Source with key '%v' and location value '%v' does not exist or is unreachable. "+
						"Please check the url or filepath and fix or consider contacting the maintainer.", k, v)
					break
				}
			}
		}
		if err != nil {
			err = nil
		} else {
			re := regexp.MustCompile("^(?:[\\w]\\:|\\.\\/)([a-z_\\-\\s0-9\\.]+(?:\\\\|\\/)?)+$")
			matched := re.MatchString(k)
			if !matched {
				err = fmt.Errorf("Invalid path for source key '%v'. Please make this a relative path in accordance "+
					"with the spec found here https://ethpm.github.io/ethpm-spec/package-spec.html#sources-sources.", k)
			}
		}
	}
	return
}

func checkContractTypes(ct map[string]*ethcontract.ContractType) (err error) {
	for k, v := range ct {
		if retErr := ethregexlib.CheckAlias(k); retErr != nil {
			err = fmt.Errorf("contract_types key '%v' does not conform to the standard. Please see "+
				"https://ethpm.github.io/ethpm-spec/glossary.html#term-contract-alias "+
				"for the spec", k)
			break
		}
		if retErr := v.Validate(k); retErr != nil {
			err = fmt.Errorf("contract_type with key '%v' returned the following error: "+
				"%v", k, retErr)
			break
		}
	}
	return
}

func checkDeployments(d map[string]map[string]*ethcontract.ContractInstance) (err error) {
OuterLoop:
	for k, v := range d {
		if retErr := ethregexlib.CheckBIP122URI(k); retErr != nil {
			err = fmt.Errorf("deployment with key '%v' returned the following error: "+
				"%v", k, retErr)
			break
		}
		for i, z := range v {
			if retErr := ethregexlib.CheckContractName(i); retErr != nil {
				err = fmt.Errorf("deployment[%v] with key '%v' returned the following error: "+
					"%v", k, i, retErr)
				break OuterLoop
			}
			dependencyLengths := make(map[string]int)
			for _, y := range z.RuntimeBytecode.LinkDependencies {
				if y.Type == "reference" {
					dependencyLengths[y.Value], err = getLinkValueDependencyLength(k, v, y.Value)
					if err != nil {
						err = fmt.Errorf("deployment[%v]:contract_instance[%v] returned the following dependency "+
							"error for dependency '%v'. Please ensure you have the dependency installed correctly: "+
							"%v", k, i, y.Value, err)
						break OuterLoop
					}
				}
			}
			if retErr := z.Validate(i, dependencyLengths); retErr != nil {
				err = fmt.Errorf("deployment[%v]:contract_instance[%v] returned the following error: "+
					"%v", k, i, retErr)
				break OuterLoop
			}
		}
	}
	return
}

func getLinkValueDependencyLength(blockchainURI string, thisDeps map[string]*ethcontract.ContractInstance, d string) (length int, err error) {
	if d == "" {
		return
	}

	var currentDir string

	if retErr := ethregexlib.CheckDependencyTree(d); retErr != nil {
		err = retErr
		return
	}
	if currentDir, err = os.Getwd(); err != nil {
		return
	}
	depTree := strings.Split(d, ":")
	treeLength := len(depTree)
	var depPath strings.Builder
	depPath.WriteString(currentDir)
	if treeLength > 1 {
		for i := 0; i < (treeLength - 2); i++ {
			depPath.WriteString("/ethpm-dependencies/" + depTree[i])
		}
		depPath.WriteString("/" + depTree[treeLength-2] + ".json")

		var file *os.File
		if file, err = os.Open(filepath.FromSlash(depPath.String())); err != nil {
			return
		}
		manifest := make([]byte, 100)
		_, err = file.Read(manifest)
		if err != nil {
			return
		}
		pm := PackageManifest{}
		err = json.Unmarshal(manifest, &pm)
		if err != nil {
			return
		}
		bc := pm.Deployments[blockchainURI][depTree[treeLength-1]].Address
		return ((len(bc) - 2) / 2), nil
	}
	bc := thisDeps[depTree[0]].Address
	return ((len(bc) - 2) / 2), nil
}

func checkBuildDependencies(bd map[string]string) (err error) {
	var uri *url.URL

	for k, v := range bd {
		uri, err = url.Parse(v)
		if err != nil {
			break
		} else {
			if a := uri.IsAbs(); !a {
				_, err = os.Stat(v)
				if err != nil {
					err = fmt.Errorf("Source with key '%v' and location value '%v' does not exist or is unreachable. "+
						"Please check the url or filepath and fix or consider contacting the maintainer.", k, v)
					break
				}
			}
		}
		re := regexp.MustCompile("^(?:[\\w]\\:|\\.\\/)([a-z_\\-\\s0-9\\.]+(?:\\\\|\\/)?)+$")
		matched := re.MatchString(k)
		if !matched {
			err = fmt.Errorf("Invalid path for source key '%v'. Please make this a relative path in accordance "+
				"with the spec found here https://ethpm.github.io/ethpm-spec/package-spec.html#build-dependencies-build-dependencies.", k)
		}
	}
	return
}
