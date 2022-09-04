package main

import "fmt"

func main() {
	/* 
		go语言的取地址符是&，放到一个变量前使用就会 返回相应变量的内存地址
		a存放在栈空间中，栈空间中存的是堆空间的地址，堆空间中存的是实际值
	*/
	var a int = 10
	fmt.Println(&a) //0xc000114000 10的堆空间

	/* 
		声明指针的格式
		var 指针变量名 *指针类型
	*/
	var ip *int
	ip = &a
	fmt.Println(ip) //0xc000114000 指向了&a 所以和&a相同
	fmt.Println(*ip) //10

	/*
		空指针 和空指针判断
		当一个指针被定义后没有分配到任何变量时，它的值为 nil
	*/
	var p *int
	fmt.Println(p) //<nil>

	if p != nil {
		fmt.Println("指针不为空")
	} else {
		fmt.Println("指针为空")
	}
}

---
package main

import "fmt"

/*
	值传递和引用传递
*/
func aaa(a int) { //此处会复制一份
	a++
}

func bbb(a *int) { //此处拿到了a的内存地址
	*a++
}

func main(){
	a := 3
	aaa(a)
	fmt.Println(a) //3
	bbb(&a)
	fmt.Println(a) //4
}

---

package main

import "fmt"

func exchange(a,b *int){ //正确的交换方式，将a 和 b栈中的值进行了互换（房子中的人）
	fmt.Println(a,b) //0xc00001a0a0 0xc00001a0a8
	fmt.Println(*a,*b) //3 4
	*a,*b = *b,*a
	fmt.Println(a,b) //0xc00001a0a0 0xc00001a0a8
	fmt.Println(*a,*b) //4 3
}

func exchange2(c,d *int){ //错误的交换方式，虽然实现了地址（房子的）互换。但是main中的指向并没有变，因为是局部变量只在该函数中生效
	fmt.Println(c,d) //0xc00001a0b0 0xc00001a0b8
	fmt.Println(*c,*d) //5 6
	c,d = d,c
	fmt.Println(c,d) //0xc00001a0b8 0xc00001a0b0
	fmt.Println(*c,*d) //6 5
}

func main(){
	a,b := 3,4
	c,d := 5,6
	exchange(&a,&b)
	exchange2(&c,&d)
	fmt.Println(&c,&d) //0xc00001a0b0 0xc00001a0b8
	fmt.Println(c,d) //5 6
}

---
package main

import "fmt"

func exchange(a,b *int)(*int,*int){
	a,b = b,a
	return a, b
}



func main(){
	a,b := 3,4
	fmt.Println(&a,&b) //a:0xc0000ac008:3 b: 0xc0000ac010:4
	c, d := exchange(&a, &b)
	fmt.Println(c,*c,d,*d) //c:b:0xc0000ac010:4 d:a:0xc0000ac008:3
	fmt.Println(&a,&b) //0xc0000ae008 0xc0000ae010
	a = *c //此时a的堆 指针地址是0xc0000ac008 内容由3变成了4
	b = *d //此时b的堆 指针地址是0xc0000ac010 内容由4不变成了4；因为b的内存地址为c010，d指针变量的内存地址为c008值为4，那么b此时也为4
	fmt.Println(a,b) //4 4

}
