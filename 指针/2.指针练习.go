package main

import "fmt"

/*指针练习
 程序定义一个int变量num的地址并打印
 将num的地址赋给指针ptr，并通过ptr去修改num的值
 */

func main(){
	num := 18
	fmt.Println(&num) //0xc0000ac008

	ptr := &num
	*ptr = 100
	fmt.Println(*ptr,num) //100 100
}
