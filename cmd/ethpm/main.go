package main

import (
	"fmt"
	"log"

	"github.com/ethpm/ethpm-go/pkg/ethpm"
)

func main() {
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
