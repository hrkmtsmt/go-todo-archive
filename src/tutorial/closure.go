package main

import "fmt"

func UpdateState() func(string) string {
		var store string = ""

		return func(nextText string) string {
				text := store
				store = nextText
				return text
		}
}

func main() {
		getText := UpdateState()
		
		fmt.Println(getText("Hello"))
		fmt.Println(getText("World"))
}
