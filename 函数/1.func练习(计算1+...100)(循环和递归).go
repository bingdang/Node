package main

import "fmt"

//循环1+2...+100=
func plus()(sum int) {
	sum = 0
	for i:=1;i<=100;i++{
		sum += i
	}
	return
}

//递归1+2...+100=
func plus2(num int)int{
	if num == 100 {
		return 100
	}
	return num + plus2(num+1)
}
//1 + (2+ (3+ (4+ (5+ ...(+100)))))

func plus3(num int)int{
	if num == 1 {
		return 1
	}
	return num + plus3(num-1)
}

//100 + (99+ (98+ (97+ (96+ ...(+1)))))

func main() {
	//fmt.Println(plus())
	fmt.Println(plus2(1))
	fmt.Println(plus3(100))
}