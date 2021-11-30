package main

import (
	"fmt"
)

func main() {
	var all, sum, max int
	fmt.Scanln(&all)
	in := make([]int, all)
	for i := range in {
		fmt.Scan(&in[i])
		sum += in[i]
		if max < in[i] {
			max = in[i]
		}
	}

	fmt.Println(sum - max)
}
