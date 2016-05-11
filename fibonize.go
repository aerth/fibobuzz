package main

import "fmt"

// Progress 1F
func genfib() func() int {
	F1 := 0
	F2 := 1
	return func() int {
		F2, F1 = F1, F1+F2
		return F2

	}
}

func main() {
	genfib := genfib()
	// Count to 45
	for i := 0; i < 45; i++ {
		fmt.Println(genfib())

	}
}
