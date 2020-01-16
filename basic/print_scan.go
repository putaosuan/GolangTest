package main

import (
	"bufio"
	"fmt"
	"os"
)

func main14 	()  {
	/*
	输入和输出
		fmt包：输入、输出
		输出：
			print 打印
			printf 格式化打印
			println 打印后换行
		格式化打印占位符：
			%v，原样输出
			%T，打印类型
			%t，布尔类型
			%s，字符串
			%f，浮点
			%d，十进制整数
			%b，二进制整数
			%o，八进制整数
			%x，%X，16进制
				%x，0-9a-f
				%X，0-9A-F
			%c，打印字符
			%p，打印地址
		输入：
			Scanln()
			Scanf()
		bufio包
	*/

	//var a int
	//var b float64
	//fmt.Println("请输入一个整数，一个浮点类型")
	//fmt.Scanln(&a,&b)//读取键盘的输入，通过操作地址，赋值给x和y 阻塞式
	//fmt.Println(a,b)
	//
	//fmt.Scanf("%d,%f",&a,&b)
	//fmt.Printf("a的数值是%d，b的数值是%f",a,b)

	fmt.Println("请输入一个字符串")
	reader:=bufio.NewReader(os.Stdin)
	s1,_:=reader.ReadString('\n')
	fmt.Println("读到的数据是",s1)
}
