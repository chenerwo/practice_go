package main

import "fmt"

func main() {
	var n [10]int
	var i int

	/* 为数组n初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for k, v := range n {
		fmt.Println("keys:", k, v)
	}

	//for j = 0; j < 10; j++ {
	//	fmt.Printf("Element[%d] = %d\n", j, n[j])
	//}
}

/**
array的注释
*/

/**
array的注释111
*/
