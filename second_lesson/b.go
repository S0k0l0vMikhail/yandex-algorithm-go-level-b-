package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	in := strings.Split(scanner.Text(), " ")

	var shop = -20
	minDistanceLeft, minDistanceRight := map[int]int{}, map[int]int{}

	for k, v := range in {
		if v == "1" {
			minDistanceLeft[k] = k - shop
		} else if v == "2" {
			shop = k
		}
	}

	shop = 20

	for i := len(in) - 1; i >= 0; i-- {
		if in[i] == "1" {
			minDistanceRight[i] = shop - i
		} else if in[i] == "2" {
			shop = i
		}
	}

	var maxDistance int
	for k, v := range in {
		if v == "1" && maxDistance < min(minDistanceLeft[k], minDistanceRight[k]) {
			maxDistance = min(minDistanceLeft[k], minDistanceRight[k])
		}
	}

	fmt.Println(maxDistance)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
