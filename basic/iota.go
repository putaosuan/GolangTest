package main

import "fmt"

func main4()  {
	const (
		a=iota
		b
		c
		d=100
		e
		f="haha"
		g
		h=iota
	)
	fmt.Println(a,b,c,d,e,f,g,h)
	//枚举中
	const(
		SPRING=iota
		SUMMER
		AUTUMN
	)
	fmt.Println(SPRING,SUMMER,AUTUMN)
}
