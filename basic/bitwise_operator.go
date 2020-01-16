package main

import "fmt"

func main12()  {
	/*
		位运算符：将数值转换位二进制后，按位操作
		按位与；&
		按位或：|
		异或^:
			二元：a^b 对应位的值不同为1，相同为0
			一元：^a
				按位取反：1--->0 0--->1
		位清空：&^
			对于a&^b
			对于b上的每个值，如果为0则取a对应位上的数值
			如果为1，则结果位就取0
		位移运算符：
			<<:按位左移，将a转为二进制，向左移动b位 a<<b
			>>:按位右移，将a转为二进制，向右移动b位 a>>b
	 */
	a:=60 //0011 1100
	b:=13 //0000 1101
	//^		0011 0001
	//&^	0011 0000
	fmt.Printf("a:%d,%b\n",a,a)
	fmt.Printf("b:%d,%b\n",b,b)

	rest1:=a&b
	fmt.Println(rest1) //12
	rest2:=a|b
	fmt.Println(rest2) //61
	rest3:=a^b
	fmt.Println(rest3) //49
	rest4:=a&^b
	fmt.Println(rest4) //48

	c:=8
	rest6:=c<<2
	fmt.Println(rest6)
	rest7:=c>>2
	fmt.Println(rest7)
}
