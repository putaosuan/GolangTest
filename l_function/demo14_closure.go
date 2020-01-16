package main

import "fmt"

func main() {
	/*
	闭包(closure):
		一个外层函数中有内层函数，该内层函数中，会操作外层函数的局部变量（外层函数中的参数，或者外层函数中直接定义的变量），
	并且该外层函数的返回值就是这个内层函数。
		这个内层函数和外层函数的局部变量，统称为闭包结构
	 */
	res1:=increment()
	fmt.Printf("%T\n",res1)
	fmt.Println(res1)
	v1:=res1()
	fmt.Println(v1)
	v2:=res1()
	fmt.Println(v2)
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())


	res2:=increment()
	fmt.Printf("%T\n",res2)
	v3:=res2()
	fmt.Println(v3)
	fmt.Println(res2())
	fmt.Println(res2())

	fmt.Println(res1())

}
func increment() func()int {//外层函数
	//定义了一个局部变量
	i:=0
	//定义了一个匿名函数，给变量自增并返回
	fun1:=func()int{//内层函数
		i++
		return i
	}
	//返回该匿名函数
	return fun1
}