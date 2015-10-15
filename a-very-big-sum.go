package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var op int32
	var sum int64 = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&op)
		sum += int64(op)
	}
	fmt.Println(sum)
}
