package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	length := n * n

	matrix := getInput(length)

	diag1 := addWithOffset(matrix, n, n+1, 0)
	diag2 := addWithOffset(matrix, n, n-1, n-1)

	diff := diag1 - diag2
	if diff < 0 {
		diff *= -1
	}

	fmt.Println(diff)

}

func getInput(length int) []int {
	// make an array of N*N length
	matrix := make([]int, length)

	var op int
	for i := 0; i < length; i++ {
		fmt.Scan(&op)
		matrix[i] = op
	}

	return matrix
}

func addWithOffset(input []int, n int, offset int, start int) int {
	var sum int = 0
	currentOffset := start
	for i := 0; i < n; i++ {
		sum += input[currentOffset]
		currentOffset += offset
	}
	return sum
}
