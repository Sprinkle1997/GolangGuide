package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello World")
	var myName string
	myName = "XiaowenChang"
	MyFirstGoFunc(myName)

	var name string

}

func MyFirstGoFunc(myName string) {
	fmt.Println(myName)
}

func MySecondGoFunc() string {
	return "return XiaowenChang"
}
