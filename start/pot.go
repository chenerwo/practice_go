package main

import (
	"fmt"
)

//pointer

func main()  {
	/*
	var cat int = 1
	var str string = "banana"

	strB := str

	fmt.Printf("%p %p %p ", &cat, &str, &strB)
	*/

	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	fmt.Printf("prt type : %T\n", ptr)
	//ptr 指针地址
	fmt.Printf("ptr pointer: %p\n", ptr)
	//对指针取值操作
	value := *ptr
	fmt.Printf("value type: %T\n", value)

	fmt.Printf("value: %s\n", value)

	v1 := *&ptr
	fmt.Printf("v1 p %p\n", v1)

	fmt.Printf("pointer line:%s", "==========\n")

	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}

//使用指针修改值

func swap(a, b *int)  {
	*a, *b = *b, *a
	//t := *a
	//*a = *b
	//*b = t
}
