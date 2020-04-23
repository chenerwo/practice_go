package main

import (
	"fmt"
)

var highRiseBuilding [30] int
var emptyBuild [25] int

var strList [] string
var numList [] int
var numEmptyList = []int{}



func main()  {
	for i := 0; i < len(highRiseBuilding); i++ {
		highRiseBuilding[i] = i + 1
	}
	fmt.Println(highRiseBuilding[10:15])

	fmt.Println(highRiseBuilding[20:])

	fmt.Println(highRiseBuilding[:2])

	fmt.Println(strList, numList, numEmptyList)

	fmt.Println(len(emptyBuild))

	fmt.Println("========================")

	//make 动态创建切片
	//make([]type, size, cap)
	as := make([]int, 2)
	bs := make([]int, 2, 10)

	fmt.Println(as, bs)

	//append 为切片追加元素
	var xs [] int
	xs = append(xs, 1, 4)
	xs = append(xs, []int{8, 9, 10}...)
	fmt.Println(xs)
	//在开头追加
	xs = append([]int{5, 6}, xs...)
	xs = append([]int{5, 6}, xs...)
	fmt.Println(xs)

	fmt.Println("========================2")
	//链式append
	var la = []int{1, 2, 3, 4, 5}
	la = append(la[:1], append([]int{22, 222, 2222}, la[1:]...)...)
	fmt.Println(la)

}
