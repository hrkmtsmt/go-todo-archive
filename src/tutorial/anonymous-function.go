package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x - y
	}

	result := add(20, 30)
	fmt.Println(result)
}

