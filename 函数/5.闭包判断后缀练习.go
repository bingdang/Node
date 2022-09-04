package main

import (
	"fmt"
	"strings"
)

//闭包练习，传入字符串，判断是否以某字串作为结尾
func SuffixJudgment(suffix string)func(string2 string)string{
	return func (name string)string{
		if !strings.HasSuffix(name,suffix){
			return name + suffix
		}else {
			return name
		}
	}
}

func main(){
	mp3suffix := SuffixJudgment(".mp3")
	fmt.Println(mp3suffix("飞驰的铂森"))
	fmt.Println(mp3suffix("飞驰的人生"))
	fmt.Println(mp3suffix("飞驰的铂森 - 平凡之路.mp3"))

	mp4suffix := SuffixJudgment(".mp4")
	fmt.Println(mp4suffix("熊比灯的铂森"))
	fmt.Println(mp4suffix("熊比灯 - 东京不太热.mp4"))
}
---
package main

import "fmt"

//返回两个闭包
func test2(base int)(func(int)int,func(int)int){
	add := func(i int)int{
		base += i
		return base
	}
	sub := func(i int)int{
		base -= i
		return base
	}
	return add,sub
}

func main(){
	f1, f2 := test2(10)
	fmt.Println(f1(1),f2(2)) //10+1=11,11-2=9
	fmt.Println(f1(3),f2(4)) //9+3=12,12-4=8
}