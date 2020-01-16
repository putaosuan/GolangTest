package main

import "fmt"

func main29() {
	/*
	数组的遍历
		方法一：arr[0],arr[1],..
		方法二：for便利
		方法三：使用range，词义："范围"
			不需要操作数组的下标，到达数组的末尾，自动结束for range循环，
			每次都获取数组中对应的下标和值
	 */
	arr:=[5]int{2,4,6,8,10}
	for index,value:=range arr{
		fmt.Printf("下标%d的值是：%d\n",index,value)
	}
	sum:=0
	for _,v:=range arr{
		sum+=v
	}
	fmt.Println("总和是：",sum)
}
