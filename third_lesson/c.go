package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 999999999 // your required line length
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	in := strings.Split(scanner.Text(), " ")

	list := make(map[string]bool)

	for _, v := range in {
		if _, ok := list[v]; !ok {
			list[v] = false
			continue
		}
		list[v] = true
	}

	for _, v := range in {
		if val := list[v]; !val {
			fmt.Print(v + " ")
		}
	}
}
