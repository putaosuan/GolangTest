package main

import "fmt"

func main36() {
	/*
	按类型来分：
		基本数据类型：int float bool string
		复合类型：array slice map struct pointer function chan
	按照特点来分：
		值类型；int float bool string array
		引用类型：slice
			传递的地址，多个变量指向了同一块内存地址
		所以，切片是引用类型的数据，存储了底层数组的引用
	 */
	s1:=[]int{1,2,3,4}
	s2:=s1
	fmt.Println(s1,s2)
	s1[1]=200
	fmt.Println(s1,s2)
	fmt.Printf("%p\n",s1)
	fmt.Printf("%p\n",s2)
	fmt.Printf("%p\n",&s1)
	fmt.Printf("%p\n",&s2)
}
