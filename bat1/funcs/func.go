package main

import "fmt"

func sum(arr []int) int {
	var sum int
	for _, num := range arr {
		if num > 10 {
			sum += 10
		} else {
			sum += num
		}
	}
	return sum
}

func main() {
	var arr = []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	var result = sum(arr)
	fmt.Println("sum of arr is: ", result)
}
