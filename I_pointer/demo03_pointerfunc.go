package main

import "fmt"

func main()  {
	/*
	函数指针
		一个指针，指向了一个函数的指针
		因为go语言中，function默认看作是一个指针，没有*

	指针函数
		一个函数，该函数的返回值是指针
	 */
	var a func()
	a=fun1
	a()

	arr1:=fun2()
	fmt.Printf("arr1的类型：%T，地址：%p，数值：%v\n",arr1,&arr1,arr1)
	p2:=fun3()
	fmt.Printf("p2的类型：%T，地址：%p，数值：%v\n",p2,&p2,p2)//这个地址打印的p2的地址
	fmt.Printf("p2指针中存储的数组的地址：%p\n",p2)

}
func fun2()[4]int  {
	arr:=[4]int{1,2,3,4}
	return arr
}
func fun3()*[4]int  {
	arr:=[4]int{5,6,7,8}
	fmt.Printf("arr的地址%p\n",&arr)
	return &arr
}
func fun1()  {
	fmt.Println("fun1...")
}
