package main

import (
	"fmt"
	"strconv" //字符转换
)

//声明变量
//及变量赋值
func declareVariable() {
	//away one
	var i int
	var str string
	var intArray [10]int
	var strArray [5]string
	var _bool bool
	var cursor *int

	//used var
	i = 1
	str = "myString char"
	_bool = true
	for j := 0; j < 10; j++ {
		intArray[j] = j
	}
	for j := 0; j < 5; j++ {
		strArray[j] = strconv.Itoa(i+10) + "Xiaowen"
	}

	//变量间交换数据值
	la := 0
	lb := 1
	la, lb = la, lb
}

func usedCursor() *iStruct {
	var _iStruct *iStruct
	_iStruct.key = 1
	_iStruct.val = "xiaowen"
	_iStruct.isValid = true
	return &iStruct
}

type iStruct struct {
	key     int
	val     string
	isValid bool
}
