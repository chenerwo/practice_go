package main

import (
	"fmt"
)

func main()  {

	const elementCount = 10

	srcData := make([]int, elementCount)

	for i := 0; i < elementCount; i++ {
		srcData[i] = i;
	}

	refData := srcData

	copyData := make([]int, elementCount)

	copy(copyData, srcData)

	srcData[0] = 999
	// 打印引用切片的第一个元素		切片不会因为等号操作进行元素的复制。
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])

	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("dd:%d\n", copyData[i])
	}

}