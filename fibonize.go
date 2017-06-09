// Fibonize simply counts the fibonacci sequence and prints to STDOUT.
package main

import (
	"fmt"
	"math/big"
	"os"
	"time"
)

var VERSION = "Undefined"

// Progress 1 F
func genfib() {
	F1 := big.NewInt(0)
	F2 := big.NewInt(1)
	var limit big.Int
	limit.Exp(big.NewInt(99), big.NewInt(999999), nil)
	for F1.Cmp(&limit) < 0 {
		F1.Add(F1, F2)
		F1, F2 = F2, F1
		fmt.Printf("%v\n", F1)
		time.Sleep(time.Millisecond * 100)
	}

}

var ever = true

func main() {

	// Show version
	if len(os.Args) > 1 {
		fmt.Println("Version #", VERSION)
		os.Exit(1)
	}

	for ever {
		genfib()
	}
}
