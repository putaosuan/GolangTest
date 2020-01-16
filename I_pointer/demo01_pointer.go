package main

import "fmt"

func main()  {
	/*
	指针：pointer
		存储了另一个变量的内存地址的变量
	 */
	//1.定义一个int类型的变量
	a:=10
	fmt.Println("a的数值是：",a)//10
	fmt.Printf("%T\n",a)//int
	fmt.Printf("a的地址是 %p\n",&a)//0xc00006e090
	//2.创建一个指针变量，用于存储变量a的地址
	var p1 *int
	fmt.Println(p1)
	p1=&a
	fmt.Println("p1中的数值是：",p1)
	fmt.Printf("p1自己的地址是%p\n",&p1)
	fmt.Println("p1的数值是a的地址，该存储的数据是",*p1)
	//3.操作指针，更改变量的数值
	*p1=200
	fmt.Println(a)
	//4.指针的指针
	var p2 **int
	fmt.Println(p2)
	p2=&p1
	fmt.Printf("%T,%T,%T\n",a,p1,p2)
	fmt.Println("p2的数值",p2)
	fmt.Printf("p2自己的地址:%p\n",&p2)
	fmt.Println("p2中存储的地址，对应的数值，就是p1的地址对应的数据，即a的地址",*p2 )
	fmt.Println(**p2)
}
