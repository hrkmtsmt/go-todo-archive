package main

import "fmt"

func sumTwoNumber(n int, m int) int {
		return n + m;
}

func main() {
		var i int = 100
		var name string = "Hiroki"

		fmt.Println(i, name)

		var t, f bool = true, false
		fmt.Println(t, f)

		var (
				j int = 200
				value string = "golang"
		)
		fmt.Println(j, value)

		var a int = 1000
		var b string
		fmt.Println(a, b)

		c := 200 + 200
		fmt.Println(c)

		fmt.Println(sumTwoNumber(a, c))
}
