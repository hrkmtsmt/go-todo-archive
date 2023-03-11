package main

import "fmt"

func main() {
		i := 0
		for {
				i++
				if i == 4 {
						break
				}
				fmt.Println("Infinity Loop")
		}

		point := 0
		for point <= 10 {
				fmt.Println(point)
				point++
		}

		for i = 0; i <= 10; i++ {
				if i % 2 == 0 {
						continue
				}
				fmt.Println(i)
		}

		array := [3]string{"Hello", "World", "!!!"}
		for i := 0; i < len(array); i++ {
				fmt.Println(array[i])
		}

		for i, v := range array {
				fmt.Println(i, v)
		}
}
