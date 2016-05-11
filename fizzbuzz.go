package main

import (
	"bufio"
	"fmt"
	"strconv"
)

import "os"

func isFizz(i int) bool {
	if i%3 == 0 {
		return true
	}
	return false
}

func isBuzz(i int) bool {
	if i%5 == 0 {
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		j := scanner.Text()
		i, err := strconv.Atoi(j)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		if isFizz(i) {
			fmt.Printf("Fizz")
			if isBuzz(i) {
				fmt.Printf("Buzz")
			}
			fmt.Printf("\n")
		} else if isBuzz(i) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
