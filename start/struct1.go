package main

/**
结构体内嵌
 */

import (
	"fmt"
)

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int // anonymous field
	innerS //anonymous field
}

func main()  {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	//=====//
	b := B{A{1, 2}, C{"aaa", "bbb"}, 3.0, 4.0}

	b.C.ax = "haha"
	fmt.Println(b)

}

type A struct {
	ax, ay int
}

type B struct {
	A
	C
	bx, by float32
}

type C struct {
	ax, bc string
}