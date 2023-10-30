package main

import "math"

// func Perimeter(width float64, height float64) float64 {
// 	return (width + height) * 2
// }

// func Area(width, height float64) float64 {
// 	return width * height
// }

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// 定义方法 func (receiverName ReceiverType) MethodName(args)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Width) * 2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

type Shape interface {
	Area() float64
}
