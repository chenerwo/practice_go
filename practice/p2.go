package main

import (
	"fmt"
)

//题目：有四个数字：1、2、3、4,能组成多个互不相同且无重复数字的三位数？各是多少？

func main() {
	num_S := []int{1, 2, 3, 4, 5}
	var num_count int
	for _, i := range num_S {
		for _, y := range num_S {
			for _, z := range num_S {
				if (i != z) && (i != y) && (y != z) {
					fmt.Println(i, y)
					num_count ++
				}
 			}
		}
	}
	fmt.Println("总数为:", num_count)
}
