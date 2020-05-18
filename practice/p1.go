package main

import (
	"fmt"
)

//切片转换为map
func slice_to_map(s_key, s_value []string) (mapObj map[string]string) {
	mapObj = map[string]string{}
	for s_key_index := range s_key {
		mapObj[s_key[s_key_index]] = s_value[s_key_index]
	}
	return
}

func main() {
	sKey := []string{"a", "b", "c", "d", "e"}
	sValue := []string{"1", "2", "3", "4", "5"}
	rMap := slice_to_map(sKey, sValue)

	fmt.Println(rMap)
}
