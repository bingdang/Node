package main

import (
	"fmt"
	"time"
)

func for1(){
	for {
		fmt.Println("懵逼中")
		time.Sleep(1*time.Second)
	}
}

func for2(){
	for i:=0;i<=5;i++{
		fmt.Println("GG")
	}
}

func for3(){
	i:=0
	for ;i<=5;{
		fmt.Println("JJ")
		i += 1
	}
}

func for4(){
	for x,y:=8,7;x>0||y>0;x,y=x-1,y-1{
		fmt.Printf("%d * %d = %d\n",x,y,x*y)
	}
}

func main() {
	for4()
}

/**
for 执行顺序
1.初始化语句
2.条件判断
3.结束语句
**/



---

package main

import "fmt"

func main() {
	for i:=1;i<5;i++{
		for j:=0;j<=5-i;j++ {
			fmt.Print(" ")
		}

		for j:=1;j<i*2;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}