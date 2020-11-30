package main

import (
	"fmt"
	// "os"
)

//!+
func max(vals ...int) (max int) {

	if len(vals) == 0 {
		fmt.Println("error: no arguments")
		return
	}

	max = vals[0]

	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) (min int) {

	if len(vals) == 0 {
		panic("error: no arguments")
	}

	min = vals[0]

	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

//!-

func main() {
	//!+main
	fmt.Println(max())           //  "0"
	fmt.Println(max(3))          //  "3"
	fmt.Println(max(1, 6, 3, -4)) //  "4"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(max(values...)) // "4"
	//!-slice

	//!+main
	fmt.Println(min())           //  "0"
	fmt.Println(min(3))          //  "3"
	fmt.Println(min(1, 6, 3, -4)) //  "4"
	//!-main

	//!+slice
	// values := []int{1, 2, 3, 4}
	fmt.Println(min(values...)) // "4"
}