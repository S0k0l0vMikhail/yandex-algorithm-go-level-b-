package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var whiteList []map[string]bool

func main() {

	var max int

	fmt.Scan(&max)

	//s := fillSlice(max)

	var in [][]string
	//var whiteList []string
	scanner := bufio.NewScanner(os.Stdin)

	var line string

	for scanner.Scan() {

		line = scanner.Text()
		if line == "HELP" {
			break
		}

		in = append(in, strings.Split(line, " "))
	}

	for i := 1; i < len(in); i += 2 {
		if in[i][0] == "YES" {
			//appendToMap(in[i-1])
			whiteList = append(whiteList, appendToMap(in[i-1]))
		}
		if in[i][0] == "NO" {
			removeFromList(in[i-1])
			//for k := range whiteList {
			//	for _, v2 := range in[i-1] {
			//		if k == v2 {
			//			whiteList[k] = false
			//		}
			//	}
			//}
		}
	}

	ans := ""
	fmt.Println(whiteList)

	//for k := range s {
	//	kStr := strconv.Itoa(k)
	//	if v, ok := whiteList[kStr]; ok && v {
	//		ans += kStr + " "
	//	}
	//}

	fmt.Println(ans)
}

func appendToMap(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}

	return m
}

func fillSlice(n int) []int {
	s := make([]int, n+1)
	for i := 0; i <= n; i++ {
		s[i] = 0
	}

	return s
}

func removeFromList(s []string) {
	for k := range whiteList {
		for _, v := range s {
			if f, ok := whiteList[k][v]; ok && f {
				whiteList[k][v] = false
			}
		}
	}
}
