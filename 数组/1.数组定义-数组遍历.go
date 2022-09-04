package main

import (
	"fmt"
)

/*
数组定义和初始化
格式： var 数组名 [数组长度]数组类型
 */

func main() {
	//1.定义数组
	var arr1 [5]int
	//2.类型推倒并赋值
	arr2 := [5]int{1,2,3,4,5}
	//3.省略长度
	arr3 := [...]int{5,4,3,2,1}
	fmt.Println(arr1,arr2,arr3) //[0 0 0 0 0] [1 2 3 4 5] [5 4 3 2 1]
	//4.指定位置赋值
	arr4 := [5]int{1,2,3}
	arr5 := [5]int{3:100,4:200}
	fmt.Println(arr4,arr5) //[1 2 3 0 0] [0 0 0 100 200]
	//5.从其他数组赋值
	arr6 := [5]int{3:100,4:200}
	var arr7[5]int
	arr7=arr6
	fmt.Println(arr7) //[0 0 0 100 200]
}

---
package main

import "fmt"

/*
 数组的遍历
 */
func main(){
	arr := [...]int{3:30,4:31}
	// 1.for循环
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}

	// 2.range取下标
	for i:=range arr{
		fmt.Println(arr[i])
	}

	//3.range取值
	for _,i:=range arr{
		fmt.Println(i)
	}
}



