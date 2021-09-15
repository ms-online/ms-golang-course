package main

import "fmt"

func main() { 
// 字符串作为key类型
	fruits := map[string]float64{
		"apple": 5.88,
		"orange": 3.65,
		"peach": 4.25,
		"watermelon": 2.56,
	}

	fmt.Println(fruits)
	fmt.Println(fruits["orange"])

	// 整数作为key类型
	courseId := map[int]string{
		12345: "html+css零基础入门",
		23456: "javascript入门到实战",
		34567: "Vue小白入门到实战",
	}

	fmt.Println(courseId)
	fmt.Println(courseId[23456])

	courseId[12345] = "html5+css3零基础入门"
	fmt.Println(courseId)

	fruits["apple"] = 6.88
	fmt.Println(fruits)
}

