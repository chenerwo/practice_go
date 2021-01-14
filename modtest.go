package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/configor"
)

func main()  {
	fmt.Println("使用外包测试:",configor.Config{})
	conOpen := sql.Open("test", "jjj");
	conOpen.Conn()
}