//函数的可见性
project01/func
package _func

import "fmt"

var A int = 100

func Afunc() {
	fmt.Println("felix")
}

---
project01/
package main

import (
	"fmt"
	_func "project01/func"
)

func main() {
	_func.Afunc()
	fmt.Println(_func.A)
}

//包内任何变量和函数都是能访问的，包外需要名字首字母大写才可以访问


//匿名函数/函数变量/defer中执行匿名函数的情况
package main

import "fmt"

//1.定义一个加函数
func add(a,b int) int {
	return a+b
}

//1.1赋值一个函数类型的变量
func test1() {
	f1 := add
	fmt.Printf("f1的类型是%T\n",f1)
	//f1的类型是func(int, int) int

	//1.2调用这个变量
	sum := f1(5,7)
	fmt.Println(sum)
	//12
}

//2.匿名函数定义和赋值写一起
func test2() {
	f1 := func(a,b int)int{
		return a+b
	}
	fmt.Printf("f1的类型是%T\n",f1)
	//f1的类型是func(int, int) int

	//2.2调用这个变量
	sum := f1(5,7)
	fmt.Println(sum)
	//12
}

//3.defer中执行匿名函数的情况
func test3(){
	var i int = 1
	//在defer中定义并触发匿名函数
	defer func() {
		fmt.Println(i)
		//100
	}()
	i = 100
	fmt.Println(i)
	//100
}

func main() {
	//_func.Afunc()
	//fmt.Println(_func.A)
	test1()
	test2()
	test3()
}
