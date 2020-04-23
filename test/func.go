package main

import (
	"fmt"
	//"math"
)

type cb func(int) int

func main()  {


	sq := func(xx int) int {
		xx += 1
		return xx
	}
	fmt.Printf("sq的值:%d\n", sq(4))

	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是2回调, x:%d\n", x)
		return x
	})

	var c1 CcOne
	c1.cc = 20.50
	fmt.Println("cc * 20的结果:", c1.getAd())
}

func testCallBack(x int, f cb)  {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是1回调x:%d\n", x)
	return x
}

/**
相当于其他语言的类声明方式 和 方法定义访问
 */
type CcOne struct {
	cc float64
}

func (c CcOne) getAd() float64 {
	return c.cc * 20
}
