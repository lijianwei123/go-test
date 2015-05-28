package main

import "fmt"

var pow = []int {1, 2, 3, 4}

func main() {
	for i, v := range pow {
		fmt.Println("%d = %d", i, v)
	}
	fmt.Println("hello world")
}
