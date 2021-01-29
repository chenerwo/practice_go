package main

import (
	"fmt"
)
//题目：一个整数，它加上100后是一个完全平方数，再加上168又是一个完全平方数，请问该数是多少

func main() {
	for i := 1; i <= 85; i ++ {
		if 168 % i == 0 {
			j := 168 / i
			if (i > j) && ((i + j) % 2 == 0) && ((i - j) % 2 == 0) {
				// m := (i + j) / 2
				n := (i - j) / 2
				x := n * n - 100
				fmt.Printf("符合条件的整数有:%d\n", x)
			}
		}
	}
}
