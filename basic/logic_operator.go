package main

import "fmt"

func main11()  {
	/*
		逻辑运算符；操作数必须是bool，运算结果也是bool
		逻辑与：&&
		逻辑或：||
	 */
	f1:=true
	f2:=false
	f3:=true
	rest1:=f1&&f2
	rest2:=f1&&f2&&f3
	fmt.Printf("rest1:%t\n",rest1)
	fmt.Printf("rest2:%t\n",rest2)
	rest3:=f1||f2
	rest4:=f1||f2||f3
	fmt.Printf("rest3:%t\n",rest3)
	fmt.Printf("rest4:%t\n",rest4)

	fmt.Printf("f1:%t,!f1:%t",f1,!f1)
}