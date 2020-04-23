package main

import (
	"fmt"
)

//将枚举值转换为字符串

//声明类型
type chipType int

const (
	None chipType = iota
	CPU
	GPU
)

func (c chipType) String() string {

	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	
	return "N/A"
}

func main()  {
	//输出cpu的值并以整型展示
	fmt.Printf("CPU`s value: %s %d\n", CPU, CPU)
}