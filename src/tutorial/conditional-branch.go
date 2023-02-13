package main

import "fmt"


func main() {
		var num int = 1
		if num == 1 {
				fmt.Println("One")
		} else {
				fmt.Println("Other number")
		}

		if str := "Hello"; str == "Hello" {
				fmt.Println("Hello World", str)
		}

		boolean := true
		if boolean := false; boolean == false {
				fmt.Println(boolean)
		}
		fmt.Println(boolean)
}
