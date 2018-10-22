package ethpm

import (
	"fmt"
	"log"
)

func ExampleGetManifestURI() {
	uri, err := GetManifestURI("0x032508890d32f30b525b78d81a2fa87f53f1433d", "test", "1.0.0", "rinkeby", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uri)
}
