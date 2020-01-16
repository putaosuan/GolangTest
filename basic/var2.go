package main

import "fmt"

func main2() {
	var num int
	num=100
	fmt.Printf("num的数值是%d, 地址是：%p\n",num ,&num)

	num=200
	fmt.Printf("num的数值是%d, 地址是：%p\n",num ,&num)

	var name string
	name="你好"
	fmt.Println(name)


	num,name,sex:=1000,"mi","nan"
	fmt.Println(num,name,sex)

	fmt.Println("--------------------")
	var m int
	fmt.Println(m)
	var n float64
	fmt.Println(n)
	var s string
	fmt.Println(s)
	var a []int//切片[]
	fmt.Println(a)
	fmt.Println(a==nil)//=空

}
