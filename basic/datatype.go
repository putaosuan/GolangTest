package main

import "fmt"

func main6()  {
	//1.布尔类型
	var b1 bool
	b1=true
	fmt.Printf("%T,%t\n",b1,b1)
	b2:=false
	fmt.Printf("%T,%t\n",b2,b2)

	//2.整数
	var i1 int8
	i1=100
	fmt.Println(i1)
	var i2 uint8 
	i2=200
	//i2=i1		cannot use i1 (type int8) as type uint8 in assignment
	fmt.Println(i2)

	var i3 int
	i3=1000
	fmt.Println(i3)
	//var i4 int64
	//语法角度：int和int64不认为是同一种类型
	//i4=i3//提示：cannot use i3 (type int) as type int64 in assignment
	//fmt.Println(i4)

	var i5  uint8
	i5=100
	var i6 byte
	i6=i5
	fmt.Println(i5,i6)
	//浮点
	var f1 float32
	var f2 float64
	f1=3.14
	f2=4.68
	fmt.Printf("%T,%.2f\n",f1,f1)
	fmt.Printf("%T,%.3f\n",f2,f2)
	fmt.Println(f1,f2)
}
