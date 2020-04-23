package main

import (
	"flag"
	"fmt"
)

//使用指针变量获取命令行的输入信息

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")

func main()  {
	flag.Parse()
	fmt.Println(*mode)

	//创建指针的另一种方法——new() 函数
	str := new(string)
	*str = "Go教程"
	fmt.Println(*str)
}


