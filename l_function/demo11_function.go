package main

import "fmt"

func main() {
	/*
	go语言的数据类型
		数值类型：整数、浮点
			进行运算操作，加减乘除、打印
		字符串：
			可以获取单个字符，截取字符串，遍历，strings包下的函数操作
		数组，切片，map...
			存储数据、修改数据、获取数据、遍历数据
		函数：
			加(),进行调用
		注意点：
			函数作为一种复合类型，可以看作是一种特殊的变量
			函数名()：将函数进行调用，函数中的代码会全部执行，然后将return的结果返回给调用处
			函数名：指向函数体的内存地址
	 */
	//函数做一个变量
	fmt.Printf("%T\n",fun1)
	fmt.Println(fun1)

	//直接定义一个函数类型的变量
	var c func(int,int)
	fmt.Printf("%T\n",c)
	fmt.Println(c)
	c=fun1
	fmt.Println(c)

	fun1(10,20)
	c(10,20)

	res1:=fun2(10,20)//只是将值赋值给res1
	res2:=fun2//将fun2的值（函数的地址）赋值给res1，res1和fun2指向同一个函数体
	fmt.Println("res1:",res1)
	fmt.Println("res2:",res2)

	fmt.Println(res2(1,2))//也可以被调用
}
func fun2(a,b int) int  {
	return a+b
}
func fun1(a,b int) {
	fmt.Printf("a:%d,b:%d\n",a,b)
}
