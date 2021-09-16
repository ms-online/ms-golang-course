package main

import "fmt"

func updateName(x string) string{
	x = "lucy"
	return x
}

func updateFruit(y map[string]float64){
	y["peach"] = 2.88
}
	// 当变量作为参数传递给函数时，go 会生成一个“副本”
	
func main() { 
  // 类型A： string，int，boolean，float，array，struct

 name := "summer"

 name = updateName(name)

 fmt.Println(name)


  // 类型B：slice，map，function
	fruits := map[string]float64{
		"apple": 5.88 ,
		"orange": 4.68,
	}

	updateFruit(fruits)
	fmt.Println(fruits)
}

