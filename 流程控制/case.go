package main

import "fmt"

func fraction(score int) string {
	g := ""
	switch  {
	case score < 0:
		g="输入错误"
		// 自动有break
		//fallthrough，可以让代码继续跑
	case score < 60:
		g="E"
	case score < 70:
		g="D"
	case score < 80:
		g="C"
	case score < 90:
		g="B"
	case score <= 100:
		g="A"
	default:
		g="A+"
	}
	return g
}

func main() {
	fmt.Println(fraction(50),
		fraction(67),
		fraction(82),
		fraction(92),
		fraction(105))
}

---

package main

import "fmt"

func main() {
	str := "a"
	switch str {
	case "a","b","c","d":
		fmt.Println("找到了")
	default:
		fmt.Println("没有")
	}
}