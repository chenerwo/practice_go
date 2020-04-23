package main

import (
	"fmt"
)

/**
创建多核多线程并发
 */
func main()  {
	for i := 0; i < 5; i++ {
		go AsyncFunc(i)
	}
}

func AsyncFunc(index int)  {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += 1
	}
	fmt.Printf("线程:%d,sum为:%d\n", index, sum)
}