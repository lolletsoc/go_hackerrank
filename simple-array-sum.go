package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	var op int
	for i := 0; i < n; i++ {
		fmt.Scan(&op)
		sum += op
	}
	fmt.Println(uint(sum))
}
