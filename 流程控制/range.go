package main

import "fmt"

func main() {
	str := "abcde"
	//只取下标
	for i:= range str {
		fmt.Printf("%d,%c\n",i,str[i])
	}

	//取下标和值
	for i,s:=range str {
		fmt.Printf("%d,%c\n",i,s)
	}
}