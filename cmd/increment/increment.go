// Command count integers
package main

import (
	"fmt"
	"math/big"
	"os"
)

var one = big.NewInt(1)
var version = ""
var builddate = ""

// 1, 2, 3, 4 ...
func increment() {
	N1 := big.NewInt(0)
	for {
		N1.Add(N1, one) // n
		fmt.Printf("%v\n", N1)
	}

}

func main() {

	// any arguments is a fatal error
	if len(os.Args) > 1 {
		println("Increment", version, builddate)
		println("Author: aerth <aerth@riseup.net>")
		println("\nerror: too many arguments")
		os.Exit(111)
	}

	for {
		increment()
	}
}
