package main

import (
	"fmt"
	"sort"
)

func main() {
	inArr := [3]int{42, 12, 18}
	inSlice := inArr[:]
	sort.Ints(inSlice)
	result := getCommonDividers(inArr[:])
	fmt.Println(result)
}

func getCommonDividers(nums []int) []int {
	divMap := make(map[int]int, nums[len(nums)-1])
	for _, val := range nums {
		for i := 2; i <= val; i++ {
			if val%i == 0 {
				divMap[i]++
			}
		}
	}
	resArr := make([]int, nums[len(nums)-1])
	resSlice := resArr[0:0]
	for key, val := range divMap {
		if val == len(nums) {
			resSlice = append(resSlice, key)
		}
	}
	return resSlice
}
