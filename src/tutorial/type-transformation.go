package main

import (
		"fmt"
		"strconv"
)

func main() {
		var i int = 1
		_float64 := float64(i)

		fmt.Printf("i = %T\n", i)
		fmt.Printf("_float64 = %T\n", _float64)

		fl := 10.5
		i2 := int(fl)
		fmt.Printf("fl = %T\n", fl)
		fmt.Printf("i2 = %T\n", i2)


		var s string = "1001"
		fmt.Printf("s = %T\n", s)

   	j, _ := strconv.Atoi(s)
		fmt.Println(j)
		fmt.Printf("j = %T\n", j)

		var k int = 2000
		sk := strconv.Itoa(k)
		fmt.Println(k)
		fmt.Printf("sk = %T\n", sk)

		var l string = "Hello World!"
		_byte := []byte(l)
		fmt.Println(_byte)

		var h string = string(_byte)
		fmt.Println(h)
}
