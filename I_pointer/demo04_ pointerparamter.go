package main

import "fmt"

func main()  {
	a:=10
	fmt.Println("调用函数前：",a)
	fun1(&a)
	fmt.Println("调用函数后：",a)
}
func fun1(p *int)  {//传递的是a的地址
	*p=100
	fmt.Println("fun1函数中，修改后的p的值",*p)
}
