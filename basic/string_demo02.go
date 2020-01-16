package main

import (
	"fmt"
	"strings"
)

func main43() {
	/*
		strings包下关于字符串的函数
	 */
	s1:="helloworld"
	//1.时候包含指定的内容
	fmt.Println(strings.Contains(s1,"abc"))
	//2.是否包含chars中任意一个字符串
	fmt.Println(strings.ContainsAny(s1,"abcd"))
	//3.统计s出现的次
	fmt.Println(strings.Count(s1,"llo"))
	//4.已xxx开头、以xxx结尾
	s2:="20190525课堂笔记.txt"
	if strings.HasPrefix(s2,"2019") {
		fmt.Println("2019年的文件")
	}
	if strings.HasSuffix(s2,".txt") {
		fmt.Println("这是一个txt文件")
	}
	fmt.Println("---------------")
	//s1:="helloworld"
	fmt.Println(strings.Index(s1,"s"))//查找substr中s的位置，存在就返回下标，不存在就返回-1
	fmt.Println(strings.IndexAny(s1,"lh"))//查找chars中任意一个字符串在s1中出现的位置
	fmt.Println(strings.LastIndex(s1,"l"))//查找substr中l在s1中最后一次出现的位置
	//字符串的拼接
	ss1:=[]string{"abc","hello","world","bcd"}
	s3:=strings.Join(ss1,"-")
	fmt.Println(s3)
	//切割
	s4:="123,abc,def,456"
	ss2:=strings.Split(s4,",")
	for i := 0; i< len(ss2);i++  {
		fmt.Println(ss2[i])
	}
	//重复，自己拼接自己count次
	s5:="hello"
	fmt.Println(strings.Repeat(s5,4))
	//替换
	s6:=strings.Replace(s1,"l","*",2)
	fmt.Println(s6)
	s7:="hello WORLD"
	fmt.Println(strings.ToLower(s7))
	fmt.Println(strings.ToUpper(s7))
	//截取
	fmt.Println(s7[:5])
	fmt.Println(s7[6:])
}