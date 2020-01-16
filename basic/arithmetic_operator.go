package main

import "fmt"

func main9()  {
	/*
		算术运算符：+,-,*,/,%,++,--
	 */
	a:=10
	b:=3
	sum:=a+b
	fmt.Printf("%d + %d=%d\n",a,b,sum)
	div:=a/b
	mod:=a%b
	fmt.Printf("%d / %d=%d\n",a,b,div)
	fmt.Printf("%d %% %d=%d\n",a,b,mod)
	c:=3
	c++
	fmt.Println(c)
}
