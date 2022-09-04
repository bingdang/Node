package main

import "fmt"

/*
 make()用来分配引用类型的内存，例如slice、map、channel，并且初始化内存
 new()用来分配各种类型的内存，但它不会初始化内存
 make()的用途不同于new()，它只能创建slice、map、channel，并返回类型为T（非指针）的已初始化（非零值）的值
*/

func main(){
	p:=new([]int)
	fmt.Println(p) //&[],分配内存不初始化
	// 10个长度，50个容量（50ml的水杯放10ml的水）
	m:=make([]int,10,50)
	fmt.Println(m) //[0 0 0 0 0 0 0 0 0 0],分配内存同时初始化

}
