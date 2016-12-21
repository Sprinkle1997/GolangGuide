package main

import (
	"fmt"
	"strconv"
	"strings"
)

//declare struct
type myStruct struct {
	key     int
	value   string
	isValid bool
}

//primery used by myStruct
//注意这里的指针，该方法有且只能由me指针使用
func (me *myStruct) SetStruct() *myStruct {
	me.key = 1
	me.value = "kelvin"
	me.isValid = true
	return &myStruct
}

func SetMyStruct() *myStruct {
	myStruct.key = 1
	myStruct.value = "Xiaowen"
	myStruct.isValid = true
	return &myStruct
}

func void() {
	fmt.Printf("")
}

func FirstLoop() {
	iLoop := make([]string, 10)
	loop := SetValueForLoop(iLoop)
}

func SetValueForLoop(loop []string) {
	for i := 0; i < 10; i++ {
		loop[i] = strconv.Itoa(i + 10) //strconv 字符转换
	}
}
