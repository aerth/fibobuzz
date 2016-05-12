package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"time"
)

// Progress 1F
func genfib() {
	F1 := big.NewInt(0)
	F2 := big.NewInt(1)
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)
	for F1.Cmp(&limit) < 0 {
		F1.Add(F1, F2)
		F1, F2 = F2, F1
		w := bufio.NewWriter(os.Stdout)
		fmt.Fprintf(w, "%v\n", F1)
		w.Flush()
		time.Sleep(time.Millisecond * 50)
	}

}
func main() {

	// Count to 45
	for {
		genfib()
	}
}
