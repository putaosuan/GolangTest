package main

import "fmt"

func main19()  {
	/*
	1.switch
		省略switch后的变量，相当于直接作用在true上
		switch{

		}
	2.case后面可以跟随多个数值
		switch 变量{
		
		}
	3.switch后面可以跟初始化语句
		switch 初始化语句;变量{
		}
	 */
	switch{
	case true:
		fmt.Println("true...")
	case false:
		fmt.Println("false...")
	}
	/*
	成绩【0-59】不及格
		【60-69】及格
		【70-79】中
		【80-89】良好
		【90-100】优秀
	 */
	score:=88
	switch  {
	case score<60:
		fmt.Println("不及格")
	case score>=60&&score<=69:
		fmt.Println("及格")
	case score>=70&&score<=79:
		fmt.Println("中")
	case score>=80&&score<=89:
		fmt.Println("良好")
	case score>=90&&score<=100:
		fmt.Println("优秀")
	default:
		fmt.Println("数据有误")
	}
	fmt.Println("------------------------")
	letter:=""
	switch letter {
	case "A","E","I","O","U":
		fmt.Println("是元音")
	case "M","N":
		fmt.Println("是M或者N")
	default:
		fmt.Println("是其他字母")
		
	}
	/*
	月份的天数
	 */
	month:=2
	day:=0
	year:=2019
	switch month {
	case 1,3,5,7,8,10,12:
		day=31
	case 4,6,9,11:
		day=30
	case 2:
		if year%400==0||year%4==0&&year%100!=0{
			day=29
		}else {
			day=28
		}
	}
	fmt.Printf("%d年%d月的天数是%d\n",year,month,day)

	fmt.Println("------------------------------")

	switch language:="golang"; language{
	case "golang":
		fmt.Println("go语言")
	case "python":
		fmt.Println("python语言")

	}
}
