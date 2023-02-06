package main

import "fmt"

func main() {
		var array [3]int
		fmt.Println(array)
		fmt.Printf("%T\n", array)

		var array2 [3]string = [3]string{"A", "B", ""}
		fmt.Println(array2)

		array3 := [3]int{1, 2, 3}
		fmt.Println(array3)

		array4 := [...]int{1, 3 ,5 ,7}
		fmt.Println(array4)
		fmt.Printf("%T\n", array4)

		fmt.Println(array4[3])

		array2[2] = "!!!!"
		fmt.Println(array2)

		fmt.Println(len(array4))
}
