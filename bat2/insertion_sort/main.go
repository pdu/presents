package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		for j := i - 1; j >= 0; j-- {
			if value < arr[j] {
				arr[j+1], arr[j] = arr[j], value
			} else {
				break
			}
		}
	}
}

func main() {
	arr := []int{6, 5, 4, 3, 2, 1, 0}
	insertionSort(arr)
	fmt.Println("sorted array: ", arr)
}
