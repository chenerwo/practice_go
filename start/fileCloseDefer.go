package main

import (
	"os"
	"fmt"
)

func main()  {
	ff := "./example.go"
	fz := fileSize(ff)
	fmt.Println(fz)
}

func fileSize(filename string) int64 {
	//根据文件名打开文件
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}

	//取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时发生错误, 关闭文件并返回文件大小为0
	if err != nil {
		f.Close()
		return 0
	}

	//取文件大小
	size := info.Size()

	return size
}

//defer 实现
func fileSizeDefer(filename string) int64 {
	f,err := os.Open(filename)

	if err != nil {
		return 0
	}

	//延迟调用close
	defer f.Close()		//不能放在第一行后，不然打开错误时f为空，

	info,err := f.Stat()

	if err != nil {
		//触发defer,
		return  0
	}
	size := info.Size()

	//return 前触发defer
	return size
}
