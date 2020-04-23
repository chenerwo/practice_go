package main

import (
	"fmt"
)

func main()  {
	var a = []int{1, 2, 3, 4, 5}

	a = a[1:]
	fmt.Println(a)

	b := a
	b = append(b[:0], b[1:]...)

	fmt.Println(b)

	c := b
	c = c[:copy(c, c[1:])]
	fmt.Println(c)
}