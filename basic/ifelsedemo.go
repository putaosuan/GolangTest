package main

import "fmt"

func main16()  {
	/*
		注意点：if后的{，一定是和if的条件写在同一行的
				else一定是if语句}之后，不能自己另起一行
	 */
	score:=0
	fmt.Println("请输入您的成绩")
	fmt.Scanln(&score)
	if score>=60 {
		fmt.Println("及格")
	}else{
		fmt.Println("不及格")
	}
}
