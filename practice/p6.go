package main

import (
	"fmt"
)

//斐波那契数列（Fibonacci sequence），又称黄金分割数列，指的是这样一个数列：0、1、1、2、3、5、8、13、21、34、……。
//使用递归方式实现输出前20个斐波那契数

func fib(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var fibs []int
	for i := 1; i <= 20; i++ {
		fibs = append(fibs, fib(i))
	}
	fmt.Println(fibs)
}
