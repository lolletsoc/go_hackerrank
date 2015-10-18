package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	positive, negative, zero := 0, 0, 0
	for i := 0; i < n; i++ {
		var op int
		fmt.Scan(&op)
		if op > 0 {
			positive++
		} else if op < 0 {
			negative++
		} else {
			zero++
		}
	}

	divider := float32(n)
	fmt.Printf("%.3f\n", float32(positive) / divider)
	fmt.Printf("%.3f\n", float32(negative) / divider)
	fmt.Printf("%.3f\n", float32(zero) / divider)
}
