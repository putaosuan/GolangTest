package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
	strconv包:字符串和基本数据类型的转换
		string convert
	 */
	//bool
	s1:="true"
	b1,err:=strconv.ParseBool(s1)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n",b1,b1)

	ss1:=strconv.FormatBool(b1)
	fmt.Printf("%T,%s\n",ss1,ss1)
	//整数
	s2:="100"
	i2,err:=strconv.ParseInt(s2,10,64)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n",i2,i2)
	fmt.Println("------------")
	ss2:=strconv.FormatInt(i2,10)
	fmt.Printf("%T,%s\n",ss2,ss2)

	//Itoa,Itoa
	i3,err:=strconv.Atoi("-42")
	fmt.Printf("%T,%d\n",i3,i3)

	ss3:=strconv.Itoa(-42)
	fmt.Printf("%T,%s\n",ss3,ss3)
}
