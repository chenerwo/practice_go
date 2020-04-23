package main

//对字符串排序

import (
	"fmt"
	"sort"
)

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main()  {
	//无序字符串
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	//使用sort排序
	sort.Strings(names)

	//遍历打印结果
	for key, v := range names {
		fmt.Printf("%d=>%s\n", key, v)
	}
}
