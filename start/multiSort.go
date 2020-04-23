package main

//对结构体的复杂排序
/*
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

//将英雄的切片定义为Heros类型
type Heros []*Hero

//实现sort.Interface接口 取元素数量方法
func (s Heros) Len() int {
	return len(s)
}

//Less方法
func (s Heros) Less(i, j int) bool {

	//如果英雄的分类不一致时，优先对分类进行排序
	if s[i].Kind < s[j].Kind {
		return s[i].Kind < s[j].Kind
	}
	return s[i].Name < s[j].Name
}

//swap
func (s Heros) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func main()  {

	//准备英雄列表
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}

	sort.Sort(heros)

	for key, v := range heros{
		fmt.Printf("%d=>%+v\n", key, v)
	}
}
 */





