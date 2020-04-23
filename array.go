package main

import "fmt"

func main() {
	var n [10]int
	var i, j int

	/* 为数组n初始化元素 */
	for i = 0; i < 11; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

/**
array的注释
*/
