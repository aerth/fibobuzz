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

// 1+1+1+1
func increment() {
	N1 := big.NewInt(0)
	N2 := big.NewInt(1)
	var limit big.Int
	limit.Exp(big.NewInt(99), big.NewInt(999999), nil)
	for N1.Cmp(&limit) < 0 {
		N1.Add(N1, N2)
		w := bufio.NewWriter(os.Stdout)
		fmt.Fprintf(w, "%v\n", N1)
		w.Flush()
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
		increment()
	}
}
