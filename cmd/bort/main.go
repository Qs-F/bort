package main

import (
	"fmt"
	"os"

	"github.com/Qs-F/bort"
)

func main() {
	b, err := bort.IsBin(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	if b {
		fmt.Println("binary")
	} else {
		fmt.Println("text")
	}
}
