package main

import (
	"fmt"
	"sort"
)

var map1 map[int]int

func main()  {
	var mapLit map[string]int

	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)

	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159

	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is : %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	fmt.Println("========================")

	//以slice 作为 map值
	//var m1 := make(map[string][]int)
	//var m2 := make(map[string]*[]int)

	//for map
	scene := make(map[string]int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for k, v := range scene{
		fmt.Printf("scene key:%s, value: %d\n", k, v)
	}

	//排序map到slice，
	var scenelist []string

	for sk := range scene{
		scenelist = append(scenelist, sk)
	}
	fmt.Println(scenelist)

	sort.Strings(scenelist)

	fmt.Println(scenelist)

	for _, sv := range scenelist{
		fmt.Printf("sort scene value: %d\n", scene[sv])
	}
}