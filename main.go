package main

import "fmt"

// 局部变量：在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。
// 全局变量：在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。
// 局部变量：形式参数也作为局部变量来使用。


var score = 100
func main() { 

 	
	sayHello("summer")
	showScore()

	for _, v := range points {
		fmt.Println(v)
	}
}

