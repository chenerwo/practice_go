package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type WebSite struct {
	Name string `xml:"name,arrt"`
	Url string
	Course []string
}

func main() {
	info := []WebSite{{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus", "http://c.biancheng.net/linux_tutorial"}},{"java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/pythos/"}}}

	//创建文件
	filePtr, err := os.Create("info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()
	//创建json编码
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}
