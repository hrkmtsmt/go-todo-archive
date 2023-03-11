package main

import "fmt"

func main() {
		var n int = 0

		switch n {
				case 1, 2:
						fmt.Println("1 or 2")
				case 0:
						fmt.Println("Zero!")
				default:
						fmt.Println("undefined")
		}

		switch i := 2; i {
				case 1, 2:
						fmt.Println("1 or 2")
				case 0:
						fmt.Println("Zero!")
				default:
						fmt.Println("undefined")
		}
		
}
