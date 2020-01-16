package main

import "fmt"

func main42() {
	//1.定义字符串
	s1:="hello"
	s2:=`hello world`
	fmt.Println(s1)
	fmt.Println(s2)
	//2.字符串的长度:返回的是字节的个数
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	//3.获取某个字节
	fmt.Println(s2[0])
	fmt.Println("-----------")
	a:='h'
	b:=104
	fmt.Printf("%c,%c,%c\n",s2[0],a,b)
	//4.字符串的遍历
	for i := 0; i< len(s2);i++  {
		//fmt.Println(s2[i])
		fmt.Printf("%c\t",s2[i])
	}
	fmt.Println()
	//for range
	for _,v:=range s2 {
		fmt.Printf("%c\t",v)
	}
	fmt.Println()
	//字符串是字节的集合
	slices:=[]byte{65,66,67,68,69}
	s3:=string(slices)
	fmt.Println(s3)

	s4:="ABCDE"
	slices2:=[]byte(s4)
	fmt.Println(slices2)
	//字符串不能修改

}
