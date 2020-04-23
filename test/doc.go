package main

import (
	"fmt"
)

const TABLE_PREFIX, VALUE_PREFIX string  = "rw_", "ys"
const UNKNOW, MALE, FEMALE  = 0, 1, 2

func main()  {
	fmt.Println("doc.go")

	v1, v2 := showConst()

	fmt.Println(v1, v2)

	area := calculatorArea(15, 20)

	fmt.Println(area)

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g, h, i, j)
}

func showConst()(string, string)  {
	return TABLE_PREFIX, VALUE_PREFIX
}

func calculatorArea(int162 int16, int163 int16) (int16) {
	return int162 * int163
}

func useSysFunc(string2 string) (int, string)  {
	length := len(string2)
	return length, string2
}

/**
--常量
常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

--iota 特殊常量
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
 */
const (
	a = iota
	b
	c
	d = "haha"
	e
	f = iota
)

const (
	g = 1<<iota
	h = 3<<iota
	i
	j
)


