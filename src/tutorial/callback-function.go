package main

import "fmt"

func CallFunc(f func()) {
		f()
}

func main() {
		CallFunc(func() {
				fmt.Println("Call!!!")
		})
}

