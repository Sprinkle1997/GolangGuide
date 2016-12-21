//Golang 中的指针
package main

import (
	"fmt"
)

//****调用

//如何使用指针
func usedCursor() {
	var i *iCursorStruct
	i = getCursor()
	fmt.Printf("%d\n", i.id) //打印指针的Id
	//遍历打印指针的data数组
	for j, val := range i.data {
		fmt.Printf("%d >> %s\t", j, val)
	}
	//output result
	/*
		运行...

		Hello World
		1
		0 >> xiaowen1	1 >> xiaowen2
		成功: 进程退出代码 0.
	*/
}

//打印data
//func UsedCursor() {
//	var j *iCursorStruct
//	for i, data := range j.data {
//		fmt.Printf("%d >> %s", i, data)
//	}
//}

///设置并获取指针的值
///
///return 指针
func getCursor() *iCursorStruct {
	_data := make([]string, 2)               //设置数据长度
	_data = []string{"xiaowen1", "xiaowen2"} //给数组赋值

	//另外一种数组定义方式
	//_data:=[1]string{"xiaowen"}
	return &iCursorStruct{id: 1, data: _data}
}

type iCursorStruct struct {
	id   int
	data []string
}
