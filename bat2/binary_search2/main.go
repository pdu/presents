package main

import "fmt"

// START OMIT
func binarySearch2(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	arr := []int{90, 91, 93, 95, 99, 100}
	fmt.Println("index of 96 is ", binarySearch2(arr, 96))
	fmt.Println("index of 95 is ", binarySearch2(arr, 95))
	fmt.Println("index of 92 is ", binarySearch2(arr, 92))
}

// END OMIT
