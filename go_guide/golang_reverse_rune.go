//内容反转
//如：xiaowen => newoaix
package main

import (
	"fmt"
)

func ReceiveVar() {
	Reverse("xiaowen")
}

func Reverse(s string) string {
	//s = "xiaowen"
	r := []rune(s) //获取字符串的每一个字符到字符数组 rune符文；符石

	//下面for循环中的变量可以如下理解
	/*
		i：=len(r)-1		声明i 并初始化值
		j:=0	声明j 并为其初始化值
		i=i+1,j=j-1
	*/
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	//尝试打印下面的内容，查看比较并理解上面的代码
	//	fmt.Println(r)         //打印内容：[110 101 119 111 97 105 120]
	//	fmt.Println(string(r)) //打印内容：newoaix
	return string(r)
}
