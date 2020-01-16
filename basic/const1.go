package main

import "fmt"

func main3()  {
	fmt.Println(100)
	fmt.Println("hello")

	const PATH string ="www.baidu.com"
	const PI=3.14
	fmt.Println(PATH)
	//fmt.Println(PI) 变量定义了要使用，常量可以不使用

	//定义一组常量
	const C1,C2,C3=100,3.14,"haha"
	fmt.Println(C1,C2,C3)
	const(
		MALE=0
		FEMALE=1
		UMKONW=3
	)
	//一组常量中，如果某个常量没有初始值，默认和上一行一致
	const(
		a int =100
		b
		c string ="ruby"
		d
		e
	)
	fmt.Printf("%T,%d\n",a,a)
	fmt.Printf("%T,%d\n",b,b)
	fmt.Printf("%T,%s\n",c,c)
	fmt.Printf("%T,%s\n",d,d)
	fmt.Printf("%T,%s\n",e,e)

	//枚举类型：使用常量组作为枚举类型。一组相关数值的数据
	const(
		SPRING=0
		SUMMER=1
		AUTUMN=2
		WINTER=3
	)
}