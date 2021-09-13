package main

import "fmt"

var someName = "someOne"

func main() { 

	// 字符串 
	// var nameOne string = "summer"
	// var nameTwo = "lucy"
	// var nameThree string 

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "Summer"
	// nameThree = "Henry"

	// fmt.Println(nameOne, nameThree)

	// nameFour := "John"

	// fmt.Println(nameFour)


	
	// 整数
	var ageOne int = 18
	var ageTwo = 20
	ageThree := 22

	fmt.Println(ageOne, ageTwo,ageThree)
	
  // https://pkg.go.dev/builtin
	// https://golang.org/ref/spec#Numeric_types
  //bits位
	var numOne int8 = 50
	var numTwo int16 = 300
	var numThree uint8 = 25

	//浮点
	var scoreOne float32 = 87.32
	var scoreTwo float64 = 3243234324.2
	scoreThree := 0.5

	fmt.Println(scoreOne, scoreTwo,scoreThree)


	fmt.Println(numOne,numTwo,numThree)

}