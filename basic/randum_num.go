package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main27() {
	/*
	生成随机数random：
		伪随机数，根据一定公式算法算出来的
		math/rand
	 */
	num1:=rand.Int()
	fmt.Println(num1)
	fmt.Println("-------------")
	for i := 0; i < 10; i++ {
		num2:=rand.Intn(10)
		fmt.Println(num2)
	}

	fmt.Println("------------")
	rand.Seed(13)
	num3:=rand.Intn(10)
	fmt.Println("-->",num3)
	//num4:=rand.Intn(10)
	//fmt.Println("-->",num4)

	t1:=time.Now()
	fmt.Println(t1)
	fmt.Printf("%T\n",t1)
	//时间戳：指定时间，距离1970年1月1日0点0分0秒，之间的时间差值：秒，纳秒
	timestamp1:=t1.Unix()
	fmt.Println(timestamp1)
	timestamp2:=t1.UnixNano()//纳秒
	fmt.Println(timestamp2)

	//step1：设置种子数，可以设置成时间戳
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num4:=rand.Intn(100)
		fmt.Println(num4)
	}
	fmt.Println("-----------")
	/*
	获取指定范围的随机数
	[13,48]
	Intn(10)获取的是[0,10)的
	 */
	num5:=rand.Intn(35)+13
	fmt.Println(num5)


}
