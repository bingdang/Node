package main

import (
	"fmt"
)

// 1.打印1到100所有的质数（大于1的自然数，除了1和它本身，不能被任何数整除）
//1.1 判断质数函数
func PrimeJudgment(base int)bool{
	if base <= 1 {
		return false
	}
	for i:=2;i<base;i++{
		if base%i==0 {
			return false
		}
	}
	return true
}
//1.2 判断指定范围内质数函数
func RangePrime(basea,basez int) {
	for i:=basea;i<=basez;i++{
		if PrimeJudgment(i) {
			fmt.Printf("%d ",i)
		}
	}
}

// 2.打印100到1000所有的水仙花数（指一个数字，各位数字立方和等于该数本身，例如153是一个水仙花数，1^3+5^3+3^3=153）
// 2.1判断水仙花函数
func NarcissusJudgment(base int)bool{
	i := base / 1000    //千位
	j := base / 100     //百位
	k := base / 10 % 10 //十位
	p := base % 10      //个位
	return i*i*i + j*j*j + k*k*k + p*p*p == base
}
// 2.2 判断指定范围内水仙花数
func RangeNarcissus(basea,basez int) {
	for i:=basea;i<=basez;i++{
		if NarcissusJudgment(i) {
			fmt.Printf("%d ",i)
		}
	}
}

// 3.传入一行字符，分别统计出其中英文字母、空格、数字和其他字符的个数 []rune()
// 3.1 判断函数
func Count(str string)(letterCount,NumberCount,SpaceCount,OtherCount int){
	runeStr:=[]rune(str)
	for i:=0;i<len(runeStr);i++{
		switch {
		case (runeStr[i] >= 'a' && runeStr[i] <= 'z') || (runeStr[i] >= 'A' && runeStr[i] <= 'Z'):
			letterCount ++
		case runeStr[i] >= '0' && runeStr[i] <= '9':
			NumberCount ++
		case runeStr[i] == ' ':
			SpaceCount ++
		default:
			OtherCount ++
		}
	}
	return
}

func main() {
	RangePrime(100,1000)
	RangeNarcissus(100,1000)
	fmt.Println(Count("felix 123321 好的"))
}

/*

101 103 107 109 113 ... 983 991 997 

153 370 371 407 

5 6 2 2

*/