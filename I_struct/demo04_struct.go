package main

import (
	"fmt"
)

func main() {
	/*
	结构体嵌套
		一个结构体中的字段，是另一个结构体类型
	 */
	b1:=Book{"Go语言",34.8}
	s1:=Studnet{"张三",18,b1}
	fmt.Println(s1)
	fmt.Println(s1.name,s1.age,s1.book.bookname,s1.book.price)
	fmt.Println("修改s1.book.bookname")
	s1.book.bookname="十万个为什么"
	fmt.Println(b1)
	fmt.Println(s1.name,s1.age,s1.book.bookname,s1.book.price)

	s2:=Studnet{"李小花",16,Book{"三国",169}}
	fmt.Println(s2)
	fmt.Println("-----------")
	s3:=Studnet2{"王五",18,&b1}
	fmt.Println(s3.name,s3.age,s3.book.bookname,s3.book.price)
	fmt.Println("修改s3.book.bookname")
	s3.book.price=2000
	fmt.Println(s3.name,s3.age,s3.book.bookname,s3.book.price)
	fmt.Println(b1)
}
//定义一个书的结构体
type Book struct {
	bookname string
	price float64
}
//定义一个学生结构体
type Studnet struct {
	name string
	age int
	book Book
}
//定义一个学生结构体
type Studnet2 struct {
	name string
	age int
	book *Book
}
