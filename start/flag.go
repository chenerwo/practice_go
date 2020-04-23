package main

import (
	"fmt"
	"flag"
)

var Input_Name = flag.String("name", "gerry", "input ur name")
var Input_Age = flag.Int("age", 20, "input ur age")
var Input_flagvar int

func Init()  {
	flag.IntVar(&Input_flagvar, "flagname", 1234, "help message for flagname")
}

func main()  {
	Init()
	flag.Parse()

	fmt.Printf("args=%s, num = %d\n", flag.Args(), flag.NArg())

	for i:=0; i < flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s", i, flag.Arg(i))
	}

	fmt.Println("name=", *Input_Name)
	fmt.Println("age=", *Input_Age)
	fmt.Println("flagname=", Input_flagvar)
}

