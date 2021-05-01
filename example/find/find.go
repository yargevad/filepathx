package main

import (
	"fmt"
	"os"

	"github.com/yargevad/filepathx"
)

func main() {
	if 2 != len(os.Args) {
		fmt.Println(len(os.Args), os.Args)
		fmt.Fprintf(os.Stderr, "Usage: go build example/find/*.go; ./find <pattern>\n")
		os.Exit(1)
		return
	}
	pattern := os.Args[1]

	matches, err := filepathx.Glob(pattern)
	if err != nil {
		panic(err)
	}

	for _, match := range matches {
		fmt.Printf("MATCH: [%v]\n", match)
	}
}
