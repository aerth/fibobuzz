// fibonize command simply prints fibonacci sequence to STANDARD OUTPUT.
package main

import (
	"fmt"
	"math/big"
	"os"
)

const (
	ever  = true
	start = 0
)

var version = ""
var builddate = ""

// 1, 1, 2, 3, 5, 8...
func genfib() {
	F1 := big.NewInt(start)
	F2 := big.NewInt(start + 1)
	for {
		F1.Add(F1, F2)

		F1, F2 = F2, F1

		// print to standard output
		fmt.Printf("%v\n", F1)
	}

}

func main() {
	// any arguments is a fatal error
	if len(os.Args) > 1 {
		println("Fibonize", version, builddate)
		println("Author: aerth <aerth@riseup.net>")
		println("\nerror: too many arguments")
		os.Exit(111)
	}

	// I'd like to see how high this can get
	for ever {
		genfib()
	}

}
