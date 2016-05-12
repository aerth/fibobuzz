package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func isPrime(i *big.Int) bool {
	j := new(big.Int)
	j.Set(i)
	if j.ProbablyPrime(30) {
		return true
	}
	return false
}

var VERSION = "undefined"
var f = big.NewInt(3)
var buzz = big.NewInt(5)
var fb = big.NewInt(15)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Version #", VERSION)
		os.Exit(0)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		b := new(big.Int)
		//x := new(big.Int)
		//y := new(big.Int)
		b.SetString(s, 0)

		//	fmt.Printf(b.Text(10))
		y := new(big.Int).Mod(b, f)
		z := new(big.Int).Mod(b, buzz)
		x := new(big.Int).Mod(b, fb)
		if x.Int64() == 0 {
			fmt.Printf("FizzBuzz")
		} else {
			if y.Int64() == 0 {
				fmt.Printf("Fizz")
			} else if z.Int64() == 0 {
				fmt.Printf("Buzz")
			} else {
				fmt.Printf(s)
			}

		}

		//	fmt.Println(x, y, z)

		//		fmt.Println(y)
		//		fmt.Println(z)
		if isPrime(b) {
			fmt.Printf(" is prime")
		}

		fmt.Printf("\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
