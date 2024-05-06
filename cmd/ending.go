package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d компьютер%s\n", i, getEnding(i))
	}
}

func getEnding(num int) string {
	remainder := num % 10

	switch remainder {
	case 0, 5, 6, 7, 8, 9:
		return "ов"
	case 2, 3, 4:
		return "a"
	}

	return ""
}
