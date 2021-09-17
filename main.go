package main

import (
	"fmt"
	"math"
)
//接口：把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

// 定义图形接口（）具有相同的方法
type shape interface {
	area() float64
	circumf() float64
}

// 定义正方形结构体
type square struct {
	length float64
}

// 定义圆结构体
type circle struct {
	radius float64
}

// 正方形的方法
// 面积
func (s square) area() float64 {
	return s.length * s.length
}
// 周长
func (s square) circumf() float64 {
	return s.length * 4
}

// 圆的方法
// 面积
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
// 周长
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

//打印图形信息
func printShapeInfo(s shape) {
	fmt.Printf("%T的面积为: %0.2f \n", s, s.area())
	fmt.Printf("%T的周长为: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 5.2},
		circle{radius: 4.8},
		circle{radius: 10.3},
		square{length: 8.6},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}
