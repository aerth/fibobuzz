package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var version = ""
var builddate = ""

const (
	fizz     = 3
	buzz     = 5
	fizzbuzz = 15
)

// small big ints ( would make const if possible ))
var f = big.NewInt(fizz)
var bz = big.NewInt(buzz)
var fb = big.NewInt(fizzbuzz)

// is probably prime ?
func isPrime(i big.Int) bool {
	if i.ProbablyPrime(30) {
		return true
	}
	return false
}

func main() {
	// any arguments is a fatal error
	if len(os.Args) > 1 {
		println("Fizzbuzz", version, builddate)
		println("Author: aerth <aerth@riseup.net>")
		println("\nerror: too many arguments, send numbers (big or small) to stdin")
		os.Exit(111)
	}

	b := new(big.Int)
	scanner := bufio.NewScanner(os.Stdin)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			for _, char := range token {
				_, hasLettersOrTooBig := strconv.ParseInt(string(char), 10, 32)
				if hasLettersOrTooBig != nil {
					println(hasLettersOrTooBig.Error())
					token = nil // just erase line and continue
				}
			}
		}
		return
	}
	scanner.Split(split)
	for scanner.Scan() {
		// read line from stdin
		s := scanner.Text()

		b.SetString(s, 0)

		// x % 15
		x := new(big.Int).Mod(b, fb)
		// x % 5
		z := new(big.Int).Mod(b, bz)
		// x % 3
		y := new(big.Int).Mod(b, f)

		switch {
		case x.Int64() == 0:
			fmt.Printf("FizzBuzz")
		case y.Int64() == 0:
			fmt.Printf("Fizz")
		case z.Int64() == 0:
			fmt.Printf("Buzz")
		default:
			fmt.Printf(s)
		}

		if isPrime(*b) {
			// a number/word has been printed, add space prefix
			fmt.Printf(" is prime")
		}

		// finally, add new line
		fmt.Printf("\n")

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
