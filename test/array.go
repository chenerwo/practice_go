package main

import "fmt"

var arr2 [15] int
var arr1 = []int{1, 2, 3, 4, 5, 6, 88, 79, 20}

var a1a2 = [3][5][7] int {
	{0, 1, 3} ,
	{11, 12, 13, 14, 15} ,
	{111, 112, 113, 114, 115, 116, 117}
}

func main()  {

	arr2[3] = 19
	fmt.Println("arr1 3的值:", arr1[3])

	var i int
	for i=0; i< 15; i++ {
		arr2[i] = i + 100
	}

	fmt.Println("a1a2 1 4 的值", a1a2[1][4])
}

