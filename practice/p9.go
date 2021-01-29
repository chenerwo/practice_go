package main

import (
	"fmt"
	"time"
)

//格式输出当前时间

func main() {
	//获取当前时间
	x := time.Now()
	fmt.Println(x.String())

	fmt.Println(x.Format("2006-01-02 15:04:05"))

	time.Sleep(time.Second)		//等待1秒
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
