package main

import "fmt"

/**
运算包括
--算数运算符
--关系运算符
--逻辑运算符
--位运算符
--赋值运算符
--其他运算符
 */

func main() {
	var a1 uint = 9
	var a2 uint = 3

	//与或非 运算
	yu, huo, fei, zy, yy := weiMath(a1, a2)
	fmt.Println(yu, huo, fei, zy, yy)

	var x int = 4
	var ptr *int
	ptr = &x
	fmt.Println("x的值:", x)
	fmt.Println("*ptr的值:", *ptr)
	fmt.Println("ptr的值:", ptr)
}

func weiMath(a1 uint, a2 uint) (uint, uint, uint, uint, uint) {
	yu := a1 & a2
	huo := a1 | a2
	fei := a1 ^ a2
	zy := a1 << a2
	yy := a1 >> a2
	return yu, huo, fei, zy, yy
}

/**
赋值运算
=,+=,-=,*=,/=,%=,
<<=	左移后赋值, >>=
&= 与等于,与后赋值
|=,^=
 */
