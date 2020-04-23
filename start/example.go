package main

import (
	"fmt"
)

//9 9乘法
func main()  {

	var a string = "aacc"
	fmt.Printf("ptr: %p", &a)
	for i := 1; i <= 9; i++ {
		for y := 1; y <= i; y++ {
			fmt.Printf("%d*%d=%d ", i, y, i*y)
		}

		fmt.Println()	//回车符
	}
}
