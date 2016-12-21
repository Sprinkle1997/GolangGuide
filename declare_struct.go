/**
xiaowen
2016-12-21
**/
package main

import (
	"fmt"
)

func structMain() {
	v := setStructValue() //给结构赋值
	//打印结构的值
	fmt.Printf("%d >> %s >> %b", v.key, v.value, v.isValid)
}

//设置xwStruct结构的值，并将其返回
func setStructValue() *xwStruct {
	return &xwcStruct{key: 1, value: "xiaowen", isValid: true}
}

type xwStruct struct {
	key     int
	value   string
	isValid bool
}

type xwcStruct struct {
	id   int
	data []string
}
