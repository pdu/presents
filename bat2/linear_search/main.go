package main

import "fmt"

func linearSearch(arr []int, target int) int {
	result := 0
	for index, value := range arr {
		result = index
		if target <= value {
			break
		}
	}
	return result
}

func main() {
	arr := []int{90, 91, 93, 95, 99, 100}
	fmt.Println("index of 96 is ", linearSearch(arr, 96))
	fmt.Println("index of 95 is ", linearSearch(arr, 95))
	fmt.Println("index of 92 is ", linearSearch(arr, 92))
}
