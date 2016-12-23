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
func (s *Square) Area() float32 {
	return s.r * s.r
}

func (c *Circle) Area() float32 {
	return c.r * 3.14 //float32 3.14->3 3.51->4
}

func getArea() {
	a := make([]GeneralContract, 2)
}
