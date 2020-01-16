package main

import "fmt"

func main() {
	/*
	结构体
		由一系列具有相同类型或者不同类型的数据构成的数据集合
		结构体成员是由一系列的成员变量构成的，这些成员变量也称为字段
	 */
	//方法一
	var p1 Person
	fmt.Println(p1)
	p1.name="王二狗"
	p1.age=23
	p1.sex="男"
	p1.address="北京市"
	fmt.Println(p1)
	//方法二
	p2:=Person{}
	p2.name="张三"
	p2.age=25
	p2.sex="男"
	p2.address="北京市"
	fmt.Println(p2)
	//方法三
	p3:=Person{name:"李小花",age:15,sex:"女",address:"深圳市"}
	fmt.Println(p3)
	p4:=Person{
		name:"隔壁老王",
		age:30,
		sex:"男",
		address:"深圳市",
	}
	fmt.Println(p4)
	//方法四
	p5:=Person{"ruby",22,"女","北京市"}
	fmt.Println(p5)
}
//定义结构体
type Person struct {
	name string
	age int
	sex string
	address string
}
