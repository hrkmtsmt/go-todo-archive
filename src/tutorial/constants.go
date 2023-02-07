package main

import "fmt"

const Pi = 3.14 // Pascal camel case is module constants

const (
		URL string = "example.com"
		SiteName string = "Example"
)

var variableMaxNumber int = 9_223_372_036_854_775_807

const BigInt = 9_223_372_036_854_775_807 + 1

const (
	i0 = iota
	i1
	i2
)

const (
	j0 = iota + 10
	j1
	j2
)

const (
		A = 1
		B
		C
		D = "A"
		E
		F
)

func main() {
		fmt.Println(Pi)
		fmt.Println(URL, SiteName)
		fmt.Println(A, B, C, D, E, F)
		fmt.Println(variableMaxNumber)
		fmt.Println(BigInt - 1)
		fmt.Println(i0, i1, i2)
		fmt.Println(j0, j1, j2)
}
