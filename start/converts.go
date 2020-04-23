package main

import (
	"fmt"
	"strconv"
)

//string 与 int 类型之间的转换

func main()  {
	//num := 100
	//str := strconv.Itoa(num)
	//fmt.Printf("type:%T value:%#v\n", str, str)

	//a to i
	str1 := "110"
	str2 := "s100"

	num1, err := strconv.Atoi(str1)

	if err != nil {
		fmt.Printf("%v 转换失败!\n", str1)
	}

	fmt.Printf("Type:%T , value : %#v\n", num1, num1)

	num2, err := strconv.Atoi(str2)

	if err != nil {
		fmt.Printf("%v 转换失败!", str2)
	} else {
		fmt.Printf("Type:%T , value : %#v\n", num2, num2)
	}
}
