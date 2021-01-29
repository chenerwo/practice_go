package main

import (
	"fmt"
	"strings"
)

func main() {
	var uids []string
	var uidstr string = "12,34,15"

	uids = strings.Split(uidstr, ",")
	fmt.Println(uids)
	nums := int32(len(uids))
	fmt.Println("用户数量", nums)
}
