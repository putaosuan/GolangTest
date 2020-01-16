	package main

import "fmt"

func main() {
	/*
	匿名结构体
		没有名字的结构体，在创建匿名结构体的同时，同时创建对象
		变量名:=struct{
			定义字段Field
		}{
			字段进行赋值
		}
	匿名字段
		默认使用数据类型作为名字，数据类型不能冲突
	 */
	s2:= struct {
		name string
		age int
	}{
		"张三",
		23,
	}
	fmt.Println(s2.name,s2.age)

	//s3:=Worker{"王二狗",18}
	//fmt.Println(s3.name,s3.age)
	s3:=Worker{"王二狗",18}
	fmt.Println(s3.string,s3.int)
}
type Worker struct {
	//name string
	//age int
	string//匿名字段，默认使用数据类型作为名字，那么数据类型不能冲突
	int
}
