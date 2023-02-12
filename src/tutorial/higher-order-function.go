package main

import "fmt"

func ReturnFunc() func() {
		return func() {
				fmt.Println("Hello World!")
		}
}

func main() {
		hello := ReturnFunc()
		hello()
}

