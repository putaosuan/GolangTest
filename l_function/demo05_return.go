package main

import "fmt"

func main()  {
	/*
	return语句：

	 */
	s:=getSum(5)
	fmt.Println(s)
	fmt.Println(getSum2())
	
	fmt.Println("-----计算周长和面积-----")
	res1,res2:=rectangle(3,4)
	fmt.Println("周长：",res1,"面积：",res2)

	res3,res4:=rectangle(4,5)
	fmt.Println("周长：",res3,"面积：",res4)
}
func rectangle(len,wid float64)(float64,float64)  {
	perimeter:=(len+wid)*2
	area:=len*wid
	return perimeter,area
}
func rectangle2(len,wid float64)(perimeter float64,area float64)  {
	perimeter=(len+wid)*2
	area=len*wid
	return
}
func getSum(n int)(int)  {
	sum:=0
	for i := 1; i <= n; i++ {
		sum+=i
	}
	return sum
}
func getSum2()(sum int)  {//定义函数时，指明要返回的数据是哪一个
	for i:=1;i<=100;i++{
		sum+=i
	}
	return
}
