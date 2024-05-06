package main

import "fmt"

func main() {
	result := getPrimeArray(11, 20)
	fmt.Println(result)
}

func getPrimeArray(min, max int) []int {
	result := make([]int, max-min)
	resSlice := result[:0]
	for i := min; i <= max; i++ {
		if isPrime(i) {
			resSlice = append(resSlice, i)
		}
	}
	return resSlice
}
func isPrime(num int) bool {
	if num <= 3 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i < num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
