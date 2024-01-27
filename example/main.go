package main

import (
	"github.com/LetsFocus/goLF/goLF"
)

func main() {
	golf := goLF.New()
	goLF.Monitor(&golf)
}
