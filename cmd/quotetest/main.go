package main

import (
	"fmt"
	"os"

	"github.com/frioux/shellquote"
	gsq "github.com/kballard/go-shellquote"
)

func main() {
	q, err := shellquote.Quote(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't quote with shellquote: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("shellquote:", q)
	fmt.Println("gsq:", gsq.Join(os.Args[1:]...))
}
