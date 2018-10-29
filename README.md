ethpm-go
=========================   

[![Join the chat at https://gitter.im/Modular-Network/Lobby](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/ethpm/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Discord](https://img.shields.io/discord/102860784329052160.svg)](https://discord.gg/crxYSF2)   

A go package which provides an [EthPM v2 package manifest](https://github.com/ethpm/ethpm-spec) reader and writer

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Layout](#layout)
- [Tools](#tools)
- [Packages](#packages)
- [Usage](#usage)
- [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# What we need   

At this point, there are 10 packages which provide enough funtionality to build an ethpm in golang. The functions defined here could even be including directly into a geth node for package management. We need contributors for the following:   

* More testing and evaluate the quality of the codebase   
* Open issues and send PR's for overall improvement   
* Build tools that make use of this repository   
* Let me know (@Hackdom on github, use the gitter link, or Hackdom#1999 if you go to Discord) if you've made a tool using this   

Finish familiarizing yourself and let us know if you have any questions!   

# Layout
This repository abides by the standard layout [as defined here](https://github.com/golang-standards/project-layout)

# Tools
This repository uses:  
* [dep for dependency management](https://golang.github.io/dep/)
* [gitflow for branch workflow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow)  

# Packages
There are ten packages defined in the `pkg` directory with the primary package being `ethpm`.   

* ethpm - https://godoc.org/github.com/ethpm/ethpm-go/pkg/ethpm   
* bytecode - https://godoc.org/github.com/ethpm/ethpm-go/pkg/bytecode   
* ethcontract - https://godoc.org/github.com/ethpm/ethpm-go/pkg/ethcontract   
* librarylink - https://godoc.org/github.com/ethpm/ethpm-go/pkg/librarylink   
* natspec - https://godoc.org/github.com/ethpm/ethpm-go/pkg/natspec   
* packageregistry - https://godoc.org/github.com/ethpm/ethpm-go/pkg/packageregistry   
* solcutils - https://godoc.org/github.com/ethpm/ethpm-go/pkg/solcutils   
* gethutils - https://godoc.org/github.com/ethpm/ethpm-go/pkg/gethutils   
* githubutils - https://godoc.org/github.com/ethpm/ethpm-go/pkg/githubutils   
* ethregexlib - https://godoc.org/github.com/ethpm/ethpm-go/pkg/ethregexlib   

# Usage
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

# Notes
This is v0.0.1 and should be treated as such. Contributions are welcome as well as any issues identified while using this code. While some of the on-chain functionality has been lightly tested, many of the full compilation, deployment, and publishing workflows have not been fully developed nor tested just yet.
