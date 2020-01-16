package main

import "fmt"

func main28() {
	/*
	数组：
		语法：1.var 数组名 [长度] 数据类型
				2.var 数组名:=[长度] 数据类型
				3.数组名:=[...] 数据类型{元素1,元素2,...}
	 */
	var num1 int
	num1=100
	fmt.Printf("%p\n",&num1)

	var arr[4] int
	arr[0]=1
	arr[1]=2
	arr[2]=3
	arr[3]=4
	fmt.Printf("%p\n",&arr[1])
	for i := 0; i < 4; i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("数组的长度：", len(arr))
	fmt.Println("数组的容量：", cap(arr))

	//数组的其他创建方式
	var a [4] int
	fmt.Println(a)

	var b =[4] int{1,2,3,4}
	fmt.Println(b)

	var c=[5] int{1,2,3}
	fmt.Println(c)

	var d =[5] int{1:1,3:4}
	fmt.Println(d)

	var e=[5] string{"张三","王二狗","李四"}
	fmt.Println(e)

	f:=[...]int{1,2,3,4,5,6}
	fmt.Println(f)
	fmt.Println(len(f))

	g:=[...]int{1:1,3:3}
	fmt.Println(g)


}
