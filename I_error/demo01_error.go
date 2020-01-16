package main

import (
	"fmt"
	"os"
)

func main1() {
	f,err:=os.Open("test.txt")
	if err!=nil{
		//log.Fatal(err)
		fmt.Println(err)
		if  ins,ok:=err.(*os.PathError);ok{
			fmt.Println("1.0p",ins.Op)
			fmt.Println("2.0p",ins.Path)
			fmt.Println("3.0p",ins.Err)
		}
		return
	}
	fmt.Println(f.Name(),"打开文件成功。。")
}
