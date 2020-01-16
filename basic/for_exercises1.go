package main

import (
	"fmt"
)

func main22()  {
	/*
	for循环练习题
	练习1：打印58-21数字
	练习2：求1-100的和
	练习3：打印100以内，能够被3整除，但不能被5整除的数字，统计被打印数字的个数，每行打印5个
	 */
	for i:= 58; i >= 21; i-- {
		fmt.Printf("%d\t",i)
	}
	fmt.Println()
	fmt.Println("-----------")
	sum:=0
	for j:=1;j<=100;j++{
		sum+=j
	}
	fmt.Printf("1-100的和是%d",sum)
	fmt.Println()
	fmt.Println("----------")
	count:=0
	for m := 0; m <= 100; m++ {

		if m%3==0 && m%5!=0 {
			//fmt.Printf("%d\t",m)
			fmt.Print(m,"\t")
			count++
			if count%5==0{
				fmt.Println()
			}
		}

	}
	fmt.Println()
	fmt.Printf("个数是：%d",count)
}
