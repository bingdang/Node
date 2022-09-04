package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "felix.txt"
	if txt, err := ioutil.ReadFile(filename);err != nil {
		fmt.Println("报错了")
	} else {
		fmt.Printf("%s\n",txt)
	}
}