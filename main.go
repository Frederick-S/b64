package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("Usage: b64 text-to-encode or b64 -d text-to-decode")

		os.Exit(1)
	}

	decode := flag.Bool("d", false, "decode")
	flag.Parse()

	if *decode {
		decoded, err := base64.StdEncoding.DecodeString(arguments[1])

		if err != nil {
			log.Fatal("Decode error")
		} else {
			fmt.Print(string(decoded))
		}
	} else {
		fmt.Print(base64.StdEncoding.EncodeToString([]byte(arguments[0])))
	}
}
