package main

import "fmt"
import "sort"

func main() {
	arr := []int{9, 7, 5, 3, 1, -1, -3, -5, -9}
	sort.Ints(arr)
	fmt.Println("sorted array: ", arr)
}
