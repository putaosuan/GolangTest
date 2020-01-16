package main

import "fmt"

func main31() {
	/*
	排序算法；
	 */
	arr:=[5] int{15,23,8,10,7}
	for i:=0;i< len(arr)-1;i++ {
		for j:=0;j< len(arr)-i-1;j++  {
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
		fmt.Println(arr)
	}
}
