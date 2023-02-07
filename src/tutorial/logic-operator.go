package main

import "fmt"


func main() {
		var t bool = true 
		var f bool = false
		fmt.Println(t && f)
		fmt.Println(t || f)
		fmt.Println(t && !f)
		fmt.Println(!!t && !f)
}
