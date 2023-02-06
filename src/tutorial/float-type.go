package main

import "fmt"

func main() {
		var _float64 float64 = 2.4;
		fmt.Println(_float64)

		fl := 3.2
		fmt.Printf("%T, %T\n", _float64, fl)

		var _float32 float32 = 3.2
		fmt.Printf("%T, %T\n", _float32, float32(_float64))

		var zero float64 = 0.0
		var posInf float64 = 1.0
		var negInf float64 = -1.0

		fmt.Println(zero / zero, posInf / zero, negInf / zero)
}
