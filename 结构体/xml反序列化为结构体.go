package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Interests struct {
	Interest []string
}

type Person struct {
	Name      string
	Age       int
	Interests Interests
}

type Final struct {
	Person []Person
}

func main() {
	var source Final
	b, err := ioutil.ReadFile("./xml/a.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	err2 := xml.Unmarshal(b, &source)
	if err != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(source)
}
---
xml/a.xml

<?xml version="1.0" encoding="UTF-8"?>
<Persons>
    <Person>
        <Name>小张</Name>
        <Age>25</Age>
        <Interests>
            <Interest>看书</Interest>
            <Interest>游泳</Interest>
        </Interests>
    </Person>
    <Person>
        <Name>小李</Name>
        <Age>32</Age>
        <Interests>
            <Interest>听音乐</Interest>
            <Interest>看电影</Interest>
        </Interests>
    </Person>
</Persons>