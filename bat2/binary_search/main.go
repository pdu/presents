package main

import "fmt"

func binarySearch1(arr []int, target int) int {
	if len(arr) == 0 {
		return 0
	}
	mid := len(arr) / 2
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binarySearch1(arr[:mid], target)
	} else {
		return binarySearch1(arr[mid+1:], target) + mid + 1
	}
}

func main() {
	arr := []int{90, 91, 93, 95, 99, 100}
	fmt.Println("index of 96 is ", binarySearch1(arr, 96))
	fmt.Println("index of 95 is ", binarySearch1(arr, 95))
	fmt.Println("index of 92 is ", binarySearch1(arr, 92))
}
