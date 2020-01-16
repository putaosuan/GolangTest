package main

import "fmt"

func main26()  {
	/*
	goto语句
	 */
	var a=10
	LOOP:
	for a<=20  {
		if a==15{
			a++
			goto LOOP
		}
		fmt.Println("a的值是：",a)
		a++
	}
	fmt.Println("----------------")
	for i := 0; i < 10; i++ {
		for j := 0; j<10 ;j++  {
			if j==2{
				goto HERE
			}
		}
	}
	//手动返回，避免执行进入标签
	return

	HERE:
	fmt.Println("done...")
}
