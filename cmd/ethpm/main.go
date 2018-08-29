package main

import (
	"fmt"
	"log"

	pkg "github.com/ethpm/ethpm-go/pkg/ethpmpackage"
)

func main() {
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
