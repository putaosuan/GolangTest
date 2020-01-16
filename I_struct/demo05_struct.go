package main

import "fmt"

func main() {
	/*
	面向对象OPP
	GO语言的结构体嵌套：
		1.模拟继承性：is-a
		type A struct{
			field
		}
		type B struct{
			A//匿名字段
			field
		}
		2.模拟聚合关系：has-a
		type A struct{
			field
		}
		type B struct{
			 a A
			field
		}
	 */
	//1.创建父类的对象
	p1:=Person{name:"张三",age:30}
	fmt.Println(p1)
	fmt.Println(p1.name,p1.age)
	//2.创建子类的对象
	s1:=Student{Person{"王二狗",18},"潭州"}
	fmt.Println(s1)
	s2:=Student{Person:Person{"王五",18},school:"北京大学"}
	fmt.Println(s2)

	var s3 Student
	s3.Person.name="李小花"
	s3.Person.age=18
	s3.school="北京大学"

	s3.name="ruby"
	s3.age=16
	fmt.Println(s3)
}
type Person struct {
	name string
	age int
}
type Student struct {
	Person
	school string
}
