package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("Usage: base64 text-to-decode")

		os.Exit(1)
	}

	text := arguments[0]

	fmt.Print(base64.StdEncoding.EncodeToString([]byte(text)))
}
