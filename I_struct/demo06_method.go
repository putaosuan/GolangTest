package main

import "fmt"

func main() {
	/*
	方法method：
		一个方法就是一个包含了接收者的函数，接收者可以是命名类型或者结构体类型的一个值或者是一个指针。
		所有给定类型的方法属于该类型的方法集
	语法
		func(接收者)方法名(参数列表)(返回值列表){
		}
	方法对比函数的不同
		A:意义
			方法：某个类别的行为功能，需要指定的接收者调用
			函数：一段独立功能的代码，可以直接调用
		B:语法
			方法：方法名可以相同，只要接收者不同就可以
			函数：命名不能冲突
	 */
	w1:=Worker{"王二狗",18}
	w1.work()

	w2:=&Worker{"Ruby",18}
	fmt.Printf("%T\n",w2)
	w2.work()

	w2.rest()
	w1.rest()

	w1.printInfo()
	c:=Cat{"green",3}
	c.printInfo()
}
//1.定义一个工人结构体
type Worker struct {
	name string
	age int
}
type Cat struct {
	color string
	age int
}
//2.定义行为方法
func (w Worker)work()  {
	fmt.Println(w.name,"在工作。。。")
}
func (w *Worker)rest()  {//w=w2;
// w=w1的地址
	fmt.Println(w.name,"在休息。。。")//结构体类型可以简写，故可以直接这样写
}
func (w *Worker)printInfo()  {
	fmt.Printf("%s,%d\n",w.name,w.age)
}
func (c *Cat)printInfo()  {
	fmt.Printf("%s,%d\n",c.color,c.age)
}

