package main

import (
	"time"
	"fmt"
)

func main() {
	const format = "03:04:05PM"

	var input string
	fmt.Scan(&input)

	t, _ := time.Parse(format, input)

	fmt.Println(t.Format("15:04:05"))
}
