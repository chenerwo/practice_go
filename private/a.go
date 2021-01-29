package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file := "./log.log"
	fileObj, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	_, err = io.WriteString(fileObj, "hello")
}
