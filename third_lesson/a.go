package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var strs []string

	for {
		line, _, err := reader.ReadLine()
		tmp := string(line) // проблема если слишком длинная строка, обрезает
		strs = append(strs, tmp)

		if err == io.EOF {
			break
		}
	}

	str1 := strs[0]
	str2 := strs[1]

	strForConv := str1
	strForRange := str2
	if len(strForConv) > len(str2) {
		strForConv = str2
		strForRange = str1
	}

	mapForCompare := strToMap(strForConv)
	ans := 0

	for _, v := range strings.Split(strForRange, " ") {
		if string(v) != " " {
			if _, ok := mapForCompare[v]; ok {
				ans++
			}
		}
	}

	fmt.Println(ans)
}

func strToMap(s string) map[string]bool {
	in := make(map[string]bool)

	for _, v := range strings.Split(s, " ") {
		if string(v) != " " {
			in[v] = true
		}
	}

	return in
}
