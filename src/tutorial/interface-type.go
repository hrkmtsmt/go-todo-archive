package main

import "fmt"

func main() {
		var x interface{} // Typescript's any????
		fmt.Println(x)

		x = 1
		fmt.Println(x)

		x = 3.14 
		fmt.Println(x)

		x = "hello" 
		fmt.Println(x)

		x = false
		fmt.Println(x)
}
