package main

import "fmt"

func main8()  {
	var a int8
	a=5
	var b int16
	b=int16(a)
	fmt.Println(b)

	f1:=3.14
	var c int
	c=int(f1)
	fmt.Println(f1,c)

	f1=float64(a)
	fmt.Println(f1)

	sum:=f1+100
	fmt.Printf("%T,%f\n",sum,sum)
}
