package main

import (
	"fmt"
	"strconv"
)

func main() {
	printTable(10)
}

func printTable(num int) {

	colWidths := make([]string, num+1)
	for i := 0; i <= num; i++ {
		maxNumInCol := i * num
		if i == 0 {
			maxNumInCol = num
		}
		colWidths[i] = "%" + strconv.Itoa(len(strconv.Itoa(maxNumInCol))+1) + "d"
	}

	for i := 0; i <= num; i++ {
		for j := 0; j <= num; j++ {
			if i == 0 && j == 0 {
				fmt.Printf("%"+strconv.Itoa(len(strconv.Itoa(num))+1)+"s", "")
				continue
			}
			if i == 0 {
				fmt.Printf(colWidths[j], j)
				continue
			}
			if j == 0 {
				fmt.Printf(colWidths[j], i)
				continue
			}
			fmt.Printf(colWidths[j], i*j)
		}
		fmt.Println()
	}

}
