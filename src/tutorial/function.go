package main

import "fmt"

func Plus(x, y int) int {
		return x + y
}

func Div(x, y int) (int, int) {
		q := x / y
		r := x % y
		return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

var x int = 10

func Update(num int) {
	x = num
}

func main() {
		i := Plus(1, 2)
		fmt.Println(i)

		j, k := Div(30, 12)
		fmt.Println(j, k)

		res := Double(200)
		fmt.Println(res)

		Update(20)
		fmt.Println(x)
}

