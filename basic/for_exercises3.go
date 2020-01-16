package main

import "fmt"

func main25() {
	/*
	打印2-100内的所有素数
	 */
	for i := 2; i<=100; i++{
		flag:=true
		for j := 2; j <  i; j++ {
			if i%j==0 {
				flag=false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}