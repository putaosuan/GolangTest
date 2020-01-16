package main

import "fmt"

func main23()  {
	/*
	题目一：
		*****
		*****
		*****
	题目二：
		1*1=1
		1*2=2 2*2=4
		...
	*/
	for i:=1;i<=3;i++  {
		for j:=1;j<=5;j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 1; i<=9;i++  {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t",j,i,i*j)
		}
		fmt.Println()
	}
}
