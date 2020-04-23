package main

//对结构体的复杂排序

import (
	"fmt"
	"sort"
)

type HeroKind int

//定义heroKind常量，类似于枚举
const (
	None1 HeroKind = iota
	Tank
	Assassin
	Mage
)

//定义英雄名单的结构
type Hero struct {
	Name string //名字
	Kind HeroKind //类型
}


func main()  {

	heros := []*Hero{
		{"吕布", Tank},
		{"李白", Assassin},
		{"妲己", Mage},
		{"貂蝉", Assassin},
		{"关羽", Tank},
		{"诸葛亮", Mage},
	}

	//使用sort.Slice 函数
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind < heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})

	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}





