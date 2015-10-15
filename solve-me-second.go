package main

import (
	"fmt"
)

func main() {
	var n int
	var op1 int
	var op2 int

	// Read the number of loops
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&op1, &op2)
		fmt.Println(uint(op1 + op2))
	}
}
