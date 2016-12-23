package main

import "fmt"

//declare type
type Square struct {
	r float32
}

type Circle struct {
	r float32
}

type CalculateArea interface {
	Area() float32 // is not func prefix
}

//inheritance
func (P T) Area() float32 {
	p.r * p.r
}

func (P T) Area() float32 {
	p.r * 3.14 //float32 3.14->3 3.51->4
}

func GetArea() {
	a := make([]GeneralContract, 2)
}
