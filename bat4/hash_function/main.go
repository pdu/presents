package main

import "fmt"

func djb2Hash(str string) int {
	ret := 5381
	for _, c := range str {
		ret = ((ret << 5) + ret) + int(c)
	}
	return ret
}

func index(str string, size int) int {
	v := djb2Hash(str) % size
	if v < 0 {
		v += size
	}
	return v
}

func main() {
	fmt.Println(djb2Hash("How are you doing?"), index("How are you doing?", 1048576))
	fmt.Println(djb2Hash("Hello world!"), index("Hello world!", 1048576))
}
