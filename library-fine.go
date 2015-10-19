package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')
	e, _ := reader.ReadString('\n')

	const timeFormat = "2 1 2006"
	actual, _ := time.Parse(timeFormat, strings.TrimSpace(a))
	expected, _ := time.Parse(timeFormat, strings.TrimSpace(e))

	sameYear := actual.Year() == expected.Year()
	sameMonth := actual.Month() == expected.Month()

	var fine int
	if actual == expected || actual.Before(expected) {
		fine = 0
	} else if sameYear && sameMonth {
		fine = 15 * (actual.Day() - expected.Day())
	} else if sameYear {
		fine = 500 * int((actual.Month() - expected.Month()))
	} else {
		fine = 10000
	}

	fmt.Println(fine)
}
