package main

import (
	"fmt"
	"github.com/Alexanderwar/cryptit/decrypt"
	"github.com/Alexanderwar/cryptit/encrypt"
)

func main() {
	fmt.Println(encrypt.Nimbus("Hello, World!"))
	fmt.Println(decrypt.Nimbus(encrypt.Nimbus("Hello, World!")))
}
