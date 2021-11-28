package main

import "fmt"

func main() {
	var all, from, to int
	fmt.Scanln(&all, &from, &to)

	if from-to == 1 || to-from == 1 {
		fmt.Println(0)
		return
	}

	if from < to {
		straight := to - from
		reverse := (from - 1) + (all - to)
		if straight > reverse {
			fmt.Println(reverse)
		} else {
			fmt.Println(straight - 1)
		}
	} else {
		straight := from - to
		reverse := (all - from) + (to - 1)

		if straight > reverse {
			fmt.Println(reverse)
		} else {
			fmt.Println(straight - 1)
		}
	}
}
