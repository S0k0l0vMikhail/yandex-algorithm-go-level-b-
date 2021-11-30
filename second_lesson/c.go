package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	in := scanner.Text()

	a := 0

	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		if in[i] != in[j] {
			a++
		}
	}

	fmt.Println(a)
}
