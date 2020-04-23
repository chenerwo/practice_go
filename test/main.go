package main

import (
	"fmt"
)

//var a bool = true;
//var b uint8 = 5
//var c string = "aabb"
//var d float32 = 33.32

var (
	type1 string
	type2 bool
)

var a, b, c uint8



func main()  {
	a = 5
	b = 8
	c = 112
	c = a
	_, j := usea()

	fmt.Println(a, b, c, j)
}

func usea()(string, int)  {
	a = 8
	return "aa", 889
}



/**
语言结构
--包声明
--引入包
--函数
--变量
--语句 & 表达式
--注释
 */