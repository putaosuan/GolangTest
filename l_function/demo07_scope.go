package main

import "fmt"

//
var num3 =100
func main() {
	/*
		作用域：
			局部变量:
			全局变量：
	*/
	fun1()
	fun2()
	fmt.Println("main中:",num3)

}
func fun1()  {
	num3=10
	fmt.Println("fun1中:",num3)
}
func fun2()  {
	fmt.Println("fun2中:",num3)
}
