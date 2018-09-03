ethpm-go
=========================
A go package which provides an [EthPM v2 package manifest](https://github.com/ethpm/ethpm-spec) reader and writer

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Layout](#layout)
- [Tools](#tools)
- [Packages](#packages)
  - [ethpm](#ethpm)
    - [Usage](#usage)
  - [natspec](#natspec)
  - [librarylink](#librarylink)
  - [ethcontract](#ethcontract)
  - [ethregexlib](#ethregexlib)
- [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Layout
This repository abides by the standard layout [as defined here](https://github.com/golang-standards/project-layout)

# Tools
This repository uses:  
* [dep for dependency management](https://golang.github.io/dep/)
* [gitflow for branch workflow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow)  

# Packages
There are six packages defined in the `pkg` directory with the primary package being `ethpm`.

## ethpm
The primary manifest object is defined in `packagemanifest.go`. `manifestinterface.go` defines a basic interface for a manifest object with a Read and Write method. We define the v2 instance which implements this interface in `v2.go`.   

### Usage
```go
package main

import (
	"fmt"
	"log"

	"github.com/ethpm/ethpm-go/pkg/ethpm"
)

func main()  {
  pm := `{"manifest_version":"2","package_name":"array-utils","version":"1.2.7"}`
  p := ethpm.PackageManifest{}

  if err := p.Read(pm); err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%+v\n", p)

  if newManifest, err := p.Write(); err != nil {
    log.Fatal(err)
  } else {
    fmt.Println(newManifest)
  }
}
```

## natspec
This package provides DevDoc, UserDoc, and DocUnion structs which correlate with natspec output

## librarylink
This provides the `LinkReference` and `LinkValue` structs which describe bytecode linking locations.

## ethcontract
This package provides `ABIObject`, which correlates with a compiler's abi output, as well as `ContractInstance` and `ContractType` which follows the EthPM v2 spec for these objects.

## ethregexlib
This package provides various regex utility functions that are relevant to ethpm and Ethereum in general.

# Notes
No testing has been implemented yet, there is no regex or any checks in the read/write functions yet either. The ethpm-spec used is contained in the api folder. You can `go run cmd/ethpm/main.go` to see the example output. Issues need to be opened to define functionality needed.
