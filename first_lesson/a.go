package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	if b == 0 && a != 0 {
		fmt.Println(3)
		return
	}
	if b == 0 && a == 0 {
		fmt.Println(c)
		return
	}
	if b == 1 {
		fmt.Println(c)
		return
	}
	if b == 4 && a != 0 {
		fmt.Println(3)
		return
	}
	if b == 4 && a == 0 {
		fmt.Println(4)
		return
	}
	if b == 6 {
		fmt.Println(0)
		return
	}
	if b == 7 {
		fmt.Println(1)
		return
	}

	fmt.Println(b)
}
