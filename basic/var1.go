package main

import "fmt"

func main1()  {
	//第一种，定义变量然后赋值
	var num1 int
	num1=30
	fmt.Printf("num1的数值是:%d\n",num1)
	//写在一行
	var num2=15
	fmt.Printf("num2的数值是%d\n",num2)
	//第二种，类型推断
	var name="张三"
	fmt.Printf("类型是：%T,数值是：%s\n",name,name)
	//第三种，简短定义，简短声明
	sum:=10
	fmt.Println(sum)
	//多个变量同时定义
	var a,b,c int
	a=1
	b=2
	c=3
	fmt.Println(a,b,c)
	var m,n int=10,100
	fmt.Println(m,n)
	var n1,m1,c1=20,3.14,"王二"
	fmt.Println(n1,m1,c1)
	var(
		studentName="李四"
		age=18
		sex="女"
	)
	fmt.Printf("姓名：%s,年龄：%d,性别：%s",studentName,age,sex)
}
