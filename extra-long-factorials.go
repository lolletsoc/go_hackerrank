package main

import (
	"fmt"
 	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(factorial(n))
}

func factorial(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(1)
	}
	temp := new(big.Int)
	temp.Mul(big.NewInt(int64(n)), factorial(n - 1))
	return temp
}
