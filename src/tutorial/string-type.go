package main

import "fmt"

func main() {
	var s string = "Hello World"
	fmt.Println(s)
	fmt.Println(s[0], string(s[0]))
	fmt.Printf("%T\n", s)

	var si string = "200"
	fmt.Println(s)
	fmt.Printf("%T\n", si)

	fmt.Println(`
		TEST
		DESU
		"Hello"
	`)
}
