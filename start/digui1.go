package main

/**
递归 互相调用  多个函数之间相互调用形成闭环
 */

import (
	"fmt"
)

func main()  {
	var count int = 6
	jie := even(count)
	fmt.Printf("count %d is even:is %t\n", count, jie)
}

func even(nr int) bool {
	fmt.Printf("even:%d\n", nr)
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr))
}

func odd(nr int) bool {
	fmt.Printf("odd:%d\n", nr)
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

