package main

import "fmt"

func main33() {
	/*
	切片slice：
		同数组类似，也叫做变长数组或者动态数组。
		特点：边长
		是一个引用类型的容器，指向了一个底层数组
		make()
		func make(t Type,size ...Integer)
		第一个参数；类型
			slice，map，chan
		第二个参数：长度len
			实际存储元素的数量
		第三个参数：容量cap
			最多能够存储的元素的数量
	append(),专门用于向切片的尾部追加元素
		slice=append(slice,元素1，元素2)
		slice=append(slice,anotherSlice)
	 */
	//数组
	arr:=[4]int{1,2,3,4}
	fmt.Println(arr)
	//切片
	var s1 [] int
	fmt.Println(s1)

	s2:=[]int{1,2,3,4}
	fmt.Println(s2)
	fmt.Printf("%T,%T\n",arr,s2)

	fmt.Println("-----------")
	s3:=make([]int,3,8)
	fmt.Println(s3)
	fmt.Printf("长度：%d，容量：%d\n", len(s3), cap(s3))

	fmt.Println("------------")
	//append
	s4:=make([]int,0,5)
	fmt.Println(s4)
	s4=append(s4,1,2,3)
	fmt.Println(s4)
	s4=append(s4,4,5)
	fmt.Println(s4)
	s4=append(s4,s3...)
	fmt.Println(s4)
}
