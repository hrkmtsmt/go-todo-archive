package main

import "fmt"

func main() {
	var i int = 100

  fmt.Println(i)
	// %T\n is format specifier
	fmt.Printf("%T\n", i)

	var i64 int64 = 64 

	fmt.Printf("%T\n", i64)

	fmt.Printf("%T\n", int32(i64))

	fmt.Println(int32(i) + int32(i64))
}

