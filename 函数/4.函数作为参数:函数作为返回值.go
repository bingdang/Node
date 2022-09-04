package main

import "fmt"

//1.1定义一个减函数
func sub(a,b int)int{
	return a - b
}

//1.2将函数作为参数传入另一个函数
func test1(a,b int,sub func(int,int)int)int{
	return sub(a,b)
}

//2.1 函数作为返回值生成闭包 (闭包是一个函数和与其相关的引用环境组合而成的实体)
func test2() func(int)int{
	var x int
	return func(d int) int {
		x += d
		return x
	}
}
/*
 局部变量x此时被匿名函数引用
 也就是x与匿名函数组成了一个实体
 所以下面函数若存活，则x会伴其左右
*/

//3.1 外部引用函数参数局部变量(不销毁)
func test3(base int) func(int)int{
	return func(d int) int {
		base += d
		return base
	}
}


func main() {
	sum := test1(9,8,sub)
	fmt.Println(sum)
	//1

	sum2 := test2()
	fmt.Println(sum2(1)) //1
	fmt.Println(sum2(10)) //11
	fmt.Println(sum2(100)) //111

	sum3 := test3(1)
	fmt.Println(sum3(1),sum3(2)) //2 4
	// 此时sum3和sum4不是一个实体了,所以不会持续加
	sum4 :=test3(17)
	fmt.Println(sum4(1),sum4(2)) //18 20
}