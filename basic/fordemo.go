package main

import "fmt"

func main21()  {
	/*
	1.标准写法
		for 表达式1;表达式2;表达式3{
		}
	2.同时省略表达式1和表达式3
	for 表达式2{
		}
	相当于while（条件）
	3.同时省略三个表达式
		for{
		}
	相当于while(true)

	 */
	i:=1
	for i<=5{
		fmt.Println(i)
		i++
	}
}
