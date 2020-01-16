package main

import "fmt"

func main() {
	/*
	数据类型
	通过指针：
		new(),不是nil，空指针
		指向了新分配的类型的内存空间，里面存储的零值
	 */
	//1.结构体是值传递
	p1:=Person{"王二狗",23,"男","北京市"}
	p2:=p1
	fmt.Println(p1)
	fmt.Println(p2)
	p2.name="张三"
	fmt.Println(p1)
	fmt.Println(p2)
	//2.定义结构体指针
	var pp1 *Person
	pp1=&p1
	fmt.Println(pp1)
	fmt.Printf("%p,%T\n",pp1,pp1)
	fmt.Println(*pp1)

	fmt.Println("---------")
	//(*pp1).name="王五"
	pp1.name="王五"//简写
	fmt.Println(*pp1)
	fmt.Println(p1)
	//使用专门的内置函数new(),go语言中专门用于创建某种类型的指针的函数
	pp2:=new(Person)
	pp2=&p1
	fmt.Printf("%T\n",pp2)
	//(*pp2).sex="女"
	pp2.sex="女"//简写
	fmt.Println(*pp2)
	fmt.Println(p1)

	pp3:=new(int)
	fmt.Println(pp3)
	fmt.Println(*pp3)
	fmt.Printf("%T\n",pp3)
}
type Person struct {
	name string
	age int
	sex string
	address string
}
