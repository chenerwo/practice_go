package main

import (
	"fmt"
)

func main()  {
	var a = complex(1, 2)
	var b = complex(3, 4)

	fmt.Println(a)
	fmt.Println(a*b)
	fmt.Println(real(a))
	fmt.Println(imag(a))
}

