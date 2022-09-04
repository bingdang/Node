package main

import "fmt"

//定义结构体
type Book struct {
	title,author string
	num,id int
}

//嵌套结构体 可外借的书
type BookLend struct {
	Book
	LendingTime string
}

//嵌套匿名结构体 不可外借的书
type BookNotLend struct {
	Book struct{
		title,author string
		num,id int
	}
	ReadTime string
}

func (t *BookLend)ChangeBookLendTime(){
	t.LendingTime = "100day"
}

func (t *BookNotLend)ChangeBookNotLendTime(){
	t.ReadTime = "100day"
}

func main() {
	book1 := &BookLend{
		Book:        Book{
			title: "Ops从入门到精通",
			author: "feichi",
			num: 101,
			id: 1,
		},
		LendingTime: "3day",
	}
	
	book2 := &BookNotLend{
		Book: struct { //匿名结构体在赋值时需要再次声明
			title, author string
			num, id       int
		}{title: "k8s从入门到精通",author: "feichi.yao",num: 200,id: 2,},
		ReadTime: "1day",
	}

	fmt.Println(book1,book2)
	book1.ChangeBookLendTime()
	book2.ChangeBookNotLendTime()
	fmt.Println(book1,book2)

}


---
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//定义临时存放json的匿名结构体
	data := &struct {
		Code int
		Msg string
	}{}
	//定义json
	jsonData := `{"code":200,"msg":""}`
	//反序列化json数据
	if err := json.Unmarshal([]byte(jsonData),data);err != nil {
		fmt.Println(err)
	}
	fmt.Println("code:",data.Code)
	fmt.Println("Msg:",data.Msg)

}