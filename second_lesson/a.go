package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var max, maxCount int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if num == 0 {
			break
		}
		if num > max {
			max = num
			maxCount = 1
		} else if num == max {
			maxCount++
		}
	}
	fmt.Println(maxCount)
}
