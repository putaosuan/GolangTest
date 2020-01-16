package main

import (
	"fmt"
)

func main() {
	/*
	函数的类型
	 */
	var a  int
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n", fun1)
	fmt.Printf("%T\n", fun2)
	fmt.Printf("%T\n", fun3)
	fmt.Printf("%T\n", fun4)
}
func fun1()  {
}
func fun2(a int)int  {
	return 0
}
func fun3(a float64,b,c int)(int,int)  {
	return 0,0
}
func fun4(a,b string ,c,d int)(string,int)  {
	return "",0
}
