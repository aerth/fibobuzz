// Fibonize simply counts the fibonacci sequence and prints to STDOUT.
package main

import (
	"bufio"
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
		w := bufio.NewWriter(os.Stdout)
		fmt.Fprintf(w, "%v\n", F1)
		w.Flush()
		time.Sleep(time.Millisecond * 300)
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
