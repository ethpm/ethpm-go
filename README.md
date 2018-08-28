ethpm-go
=========================
A go package which provides an EthPM v2 package manifest reader and writer

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Layout](#layout)
- [Tools](#tools)
- [Packages](#packages)
  - [ethpmpackage](#ethpmpackage)
    - [Usage](#usage)
  - [natspec](#natspec)
  - [ethabi](#ethabi)
  - [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Layout
This repository abides by the standard layout [as defined here](https://github.com/golang-standards/project-layout)

# Tools
This repository uses [dep for dependency management](https://golang.github.io/dep/)

# Packages
There are three packages defined in the `pkg` directory with the primary package being `manifest`.

## ethpmpackage
The primary manifest object is defined in packagemanifest.go. manifestinterface.go defines a basic interface for a manifest object with a Read and Write method. We define the v2 instance which implements this interface in v2.go.   

### Usage
```go
package main

import (
  pkg "github.com/ethpm/ethpm-go/pkg/ethpmpackage"
  "fmt"
  "log"
)

func main()  {
  pm := `{"manifest_version":"2","package_name":"ArrayUtils","version":"1.2.7"}`
  p := pkg.PackageManifest{}

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

## ethabi
This package provides an ABIObject which correlates with a compiler's abi output

## Notes
No testing has been implemented yet, there is no regex or any checks in the read/write functions yet either. The ethpm-spec used is contained in the api folder. You can `go run main.go` in the `cmd/ethpm` folder to see the example output. Issues need to be opened to define functionality needed.
