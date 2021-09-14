package main

import "fmt"

func main() { 
	age := 25
	name := "summer"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("Hi! \n")
  
	// Println
	fmt.Println("你好！")
	fmt.Println("我是Summer.")
	fmt.Println("我的名字叫",name, ",今年",age,"周岁。")

	// Printf : 格式化输出函数
   fmt.Printf("我的名字叫%v,今年%v周岁。\n", name, age)
   fmt.Printf("我的名字叫%q,今年%q周岁。\n", name, age)
   fmt.Printf("age变量的类型为%T \n",age)
	 fmt.Printf("你的得分为%0.1f! \n",95.56)

	 //Sprintf: 格式化并返回一个字符串（不带输出）
	var str =  fmt.Sprintf("我的名字叫%v,今年%v周岁。\n", name, age)
	fmt.Print(str)

	// 格式化符号：https://pkg.go.dev/fmt@go1.17.1

}