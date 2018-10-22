package ethpm

import "github.com/ethpm/ethpm-go/pkg/ethcontract"

// ManifestInterface The interface for an ethpm PackageManifest type
type ManifestInterface interface {
	Read(s string) (err error)
	Write() (s string, err error)
	WriteToDisk(directoryname string) (err error)
	AddDependency(name string, uri string)
	AddContractType(compiler string, settingsjsonstring string, compileroutputjson string, contractname string) (err error)
	AddDeployment(blockchainuri string, d *ethcontract.DeployedContractInfo)
	SourceInliner(contractdir string, sourcerelativepath string, sourcetype string) (err error)
	AddLocalPathForSource(contractdir string, sourcerelativepath string, sourcetype string) (err error)
	CompileAndValidateSource(compiler string,
		projectdir string,
		contractname string,
		inline bool,
		filepath string,
		optimize bool,
		runs int,
	) (valid bool, producedobject string, err error)
	PublishToRepositoryWithPassword(repositoryaddressashex string,
		manifesturi string,
		fromaddressashex string,
		gaspriceinwei int64,
		chainname string,
		gethdatadir string,
	) (err error)

	Validate() (err error)
}
