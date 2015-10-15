package main
import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get numbers
	v1, _ := reader.ReadString('\n')
	v2, _ := reader.ReadString('\n')

	v1 = strings.TrimSpace(v1)
	v2 = strings.TrimSpace(v2)

	n1, _ := strconv.Atoi(v1)
	n2, _ := strconv.Atoi(v2)

	fmt.Println(n1 + n2)
}
