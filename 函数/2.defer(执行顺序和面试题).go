package main

import "fmt"

func deff() {
	defer fmt.Println("最后执行")
	fmt.Println("正常语句")
}
//正常语句
//最后执行

func deffs() {
	defer fmt.Println("第1道门")
	defer fmt.Println("第2道门")
	defer fmt.Println("第3道门")
}
//先写的后执行
//第3道门
//第2道门
//第1道门

//不管出不出异常肯定执行，一般用于关闭资源

//面试题1
func deff1() {
	for i:=0;i<5;i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
//先写的后执行
//done
//4
//3
//2
//1
//0

//面试题2
func deff2() {
	var i int =10
	//此时defer已经为10，只不过暂且不执行而已
	defer fmt.Println(i)

	i = 1000
	fmt.Println(i)
}
//1000
//10


func main() {
	//deffs()
	//deff1()
	deff2()
}
