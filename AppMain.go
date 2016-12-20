package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello World")
	var myName string
	myName = "XiaowenChang"
	MyFirstGoFunc(myName)
}

func MyFirstGoFunc(myName string) {
	fmt.Println(myName)
}
