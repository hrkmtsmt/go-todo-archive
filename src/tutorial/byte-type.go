package main

import "fmt"

func main() {
		byteA := []byte{72, 73}
		fmt.Println(byteA)

		fmt.Println(string(byteA))

		_slice := []byte("My name is Hiroki!")
		fmt.Println(_slice)
		fmt.Println(string(_slice))
}
