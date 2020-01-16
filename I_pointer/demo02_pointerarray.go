package main

import "fmt"

func main()  {
	/*
	数组指针：首先是一个指针，一个数组的地址
		*[4]type
	指针数组：首先是一个数组，存储的数据类型是指针
		[4]*type

		*[5]float64,指针，一个存储了5个浮点类型数据的数组的指针
		[3]*[string],数组，存储了3个字符串的指针地址的数组
		*[5]*float64,指针，存储了5个float类型的数据的指针地址的数组的指针
		*[3]*string,存储了3个字符串的指针地址的数组的指针
		**[4]string,指针，存储了4个字符串数据的数组的指针的指针
		**[4]*string，指针，存储了4个字符串数据的指针地址的数组，的指针的指针
	 */
	//1.创建一个普通的数组
	arr1:=[4]int{1,2,3,4}
	fmt.Println(arr1)
	//2.创建一个指针，存储该数组的地址--->数组指针
	var p1 *[4]int
	p1=&arr1
	fmt.Println(p1)
	fmt.Printf("%p\n",p1)//p1存储的数组a的地址
	fmt.Printf("%p\n",&p1)//p1自己的地址

	//3.根据数组指针，操作数组
	(*p1)[0]=100
	fmt.Println(arr1)

	p1[1]=200//简化写法
	fmt.Println(arr1)

	fmt.Println("--------------")
	//4.指针数组
	a:=1
	b:=2
	c:=3
	d:=4
	arr2:=[4]int{a,b,c,d}
	arr3:=[4]*int{&a,&b,&c,&d}
	fmt.Println(arr2)
	fmt.Println(arr3)
	arr2[0]=100
	fmt.Println(a)
	fmt.Println(arr2)

	*arr3[1]=200
	fmt.Println(b)
	fmt.Println(arr2)//说明传递过来的是值

	fmt.Println("---------------")
	c=1000
	fmt.Println(arr2)
	for i := 0; i< len(arr3);i++  {
		fmt.Println(*arr3[i])
	}

}
