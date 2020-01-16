package main

import "fmt"

func main() {
	/*
	defer的语义："延迟","推迟"
	在go语言中，使用defer关键字来延迟一个函数或者方法的执行
	1.defer函数或者方法：一个函数或者方法的执行被延迟了
	2.defer的使用
		A：对象.close()，临时文件的删除
			文件.open()
			defer close()
			读或者写
		B:	go语言中关于异常的处理，使用panic()或者recover()
			panic函数用于引发恐慌，导致程序中断执行
			recover用于恢复程序的执行，recover()语法上要求必须在defer中执行
	3.如果多个defer函数：
		先延迟的后执行，后延迟的先执行
	4.defer函数传递参数的时候：
		defer函数调用时，就已经传递参数了，只是暂时不执行函数中的代码而已
	5.defer函数的注意点：
		当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的执行结束
		当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正的返回
		当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕后，该运行时恐慌才会真正被扩展至调用函数
	 */
	fmt.Println("张三")
	defer fun1("hello")
	fmt.Println("王二狗")
	defer fun1("world")

	fmt.Println("-------------")
	a:=2
	fmt.Println(a)
	defer fun2(a)
	a++
	fmt.Println("main中的值：",a)

	fmt.Println(fun3())
}
func fun3()int  {
	fmt.Println("fun3执行。。。")
	defer fun1("haha")
	return 0
}
func fun2(i int)  {
	fmt.Println(i)
}
func fun1(s string)  {
	fmt.Println(s)
}
