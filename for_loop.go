package main

import (
	"fmt"
	"strconv"
)

func forMain() {
	var arr []int
	var arrs []string
	fmt.Println("arr Array:")
	arr = setIntArray(arr)
	fmt.Println("arrs Array:")
	arrs = setStringArray(arrs)
	for i, v := range arrs {
		fmt.Printf("\t%d >> %s\n", i, v)
	}

	//output
	/*
		运行...

		arr Array:
			0 >> 5
			1 >> 5
			2 >> 6
			3 >> 6
			4 >> 7
			5 >> 7
			6 >> 8
			7 >> 8
			8 >> 9
			9 >> 9
		arrs Array:
			0 >> Hello the 0 xiaowen
			1 >> Hello the 1 xiaowen
			2 >> Hello the 2 xiaowen
			3 >> Hello the 3 xiaowen
			4 >> Hello the 4 xiaowen
			5 >> Hello the 5 xiaowen
			6 >> Hello the 6 xiaowen
			7 >> Hello the 7 xiaowen
			8 >> Hello the 8 xiaowen
			9 >> Hello the 9 xiaowen

		成功: 进程退出代码 0.
	*/
}

//这里需要注意
//传入的数据长度为0
//这里需要注意
//传入的数据长度为0
func setIntArray(arr []int) []int {
	_arr := make([]int, 10)
	arr = _arr
	for i := 0; i < 10; i++ {
		_arr[i] = (i + 1*10) / 2
		fmt.Printf("\t%d >> %d\n", i, arr[i])
	}
	return arr
}

func setStringArray(arrs []string) []string {
	_arrs := make([]string, 10)
	arrs = _arrs
	for i := 0; i < 10; i++ {
		arrs[i] = "Hello the " + strconv.Itoa(i) + " xiaowen"
	}
	return arrs
}
