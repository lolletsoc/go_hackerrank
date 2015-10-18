package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		spaces := strings.Repeat(" ", n-i)
		stairs := strings.Repeat("#", i)
		fmt.Println(spaces + stairs)
	}
}
