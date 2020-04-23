package main

import "fmt"

//func main()  {
//	var a int = 20
//	var ip *int
//
//	ip = &a;
//
//	fmt.Printf("a的变量地址为%x\n", &a);
//	fmt.Printf("ip的变量储存的指针地址为%x\n", ip)
//	fmt.Printf("ip指针的值%d", *ip);
//}

func main()  {
	var ptr *int

	if (ptr != nil) {
		fmt.Printf("不为nil")
	} else {
		fmt.Printf("为nil")
	}
	//fmt.Printf("ptr 的值为:%x\n", ptr)
}