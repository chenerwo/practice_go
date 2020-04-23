package main

import (
	"fmt"
)

var scene map[string]int

func main()  {
	scene = make(map[string]int)

	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	delete(scene, "china")
	fmt.Println(scene)
	//清空所有的方法，重新make
	scene = make(map[string]int)
	fmt.Println(scene)

}