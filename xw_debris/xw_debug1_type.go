package main

import "fmt"

type Square struct{ r, s float64 }
type Circle struct{ r float64 }

func DebugOneMain() {
	s := Square{3.14, 2.8}     //init type value
	r := Circle{3.999999999}   //init type value
	fmt.Println(s, r, s.r*r.r) //print result
}
