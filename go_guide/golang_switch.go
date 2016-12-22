package main

import (
	"fmt"
)

type xwEmptyCtr int

/*
!!!该文件Code未调试
*/
func xwUsedSwitch() {
	var res int
	res = xwSwitchFunc()
	fmt.Printf("swithc return result : %d ", res)
}

//使用Switch选择结构
func (e *xwEmptyCtr) xwSwitchFunc() int {
	switch e {
	case xwObjOne:
		return 1
	case xwObjTwo:
		return 2
	}
	return -1
}

var (
	xwObjOne = new(xwEmptyCtr) //int 类型对象
	xwObjTwo = new(xwEmptyCtr) //上面定义的int类型对象
)
