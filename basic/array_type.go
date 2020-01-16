package main

import "fmt"

func main30() {
	/*
	值类型：理解为存储的数值本身
		将数据传递给其他变量，传递的是数据的副本（备份）
		int,float,bool,string,array
	引用类型：理解为存储的数据的内存地址
		slice,map,...
	 */
	a:=3
	b:=4
	arr1:=[4] int{1,2,3,4}
	arr2:=arr1
	fmt.Println(a==b)//比较的是数值
	fmt.Println(arr1==arr2)//比较的是对应位置的数值
}
