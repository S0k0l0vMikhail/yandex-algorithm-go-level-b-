package first_lesson

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var all int

	fmt.Scan(&all)
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 10000000  // your required line length
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()

	in := strings.Split(scanner.Text(), " ")


	a := all/2

	fmt.Println(in[a])
}