package first_lesson

import "fmt"

func main() {
	var a, b, year int
	fmt.Scanln(&a, &b,&year)

	if a > 12 || b > 12 {
		fmt.Println(1)
		return
	}

	if a <= 12 && b <= 12 && a == b {
		fmt.Println(1)
		return
	}

	fmt.Println(0)
}
