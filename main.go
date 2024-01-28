package main

import (
	"fmt"
	"github.com/LetsFocus/goLF/goLF"
	"net/http"
)

func main() {
	golf := goLF.New()

	fmt.Println(golf)

	http.ListenAndServe(":8080", nil)
}
