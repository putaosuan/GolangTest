package main

import "fmt"

func main7()  {
	var s1 string
	s1="张三"
	fmt.Printf("%T,%s\n",s1,s1)
	s2:=`hello world`
	fmt.Printf("%T,%s\n",s2,s2)
	//区别""和''
	v1:='A'
	v2:="A"
	fmt.Printf("%T,%d\n",v1,v1)
	fmt.Printf("%T,%s\n",v2,v2)
	v3:='中'
	fmt.Printf("%T,%d %c %q\n",v3,v3,v3,v3)
	//转义字符
	//A:有一些字符，有特殊作用，可以转义为普通字符
	//B:有一些字符，就是一个普通字符，转义后有特殊作用 \n,换行 \t,制表符
	fmt.Println("\"hello world\"")
	fmt.Println("\"hello\nworld\"")
	fmt.Println(`hell"owor"ld`)
	fmt.Println("hell`owor`ld")
}
