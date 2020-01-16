package main

import "fmt"

func main()  {
	/*
	参数的传递：
		值传递
		引用传递
	 */
	arr1:=[4]int{1,2,3,4}
	fmt.Println("执行前：",arr1)
	fun11(arr1)
	fmt.Println("执行后：",arr1)
}
func fun11(arr2 [4]int)  {
	fmt.Println("arr2执行前：",arr2)
	arr2[0]=100
	fmt.Println("arr2执行后：",arr2)
}
