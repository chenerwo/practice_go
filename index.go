package main

import "fmt"

func main()  {
	fmt.Println("Hello,World")
}

func init()  {
	fmt.Println("Hello,World," + "init")
	//fmt.Println(arr);
}

/**
25常用关键字：
break,default,func,interface,select,case,defer,go,map,struct,chan,else,goto,
package,switch,const,fallthrough,if,range,type,continue,for,import,return,var
36预定义标识符
append,bool,byte,cap,close,complex,complex63,complex128,unit16
copy,false,float32,float64,imag,int,int8,int16,unit32
int32,int64,iota,len,make,new,nil,panic,unit64,
print,println,real,recover,string,true,unit,unit8,unitptr

程序一般由关键字、常量、变量、运算符、类型和函数组成。.、,、
*/

