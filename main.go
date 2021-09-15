package main

import "fmt"

func main() { 
	// x := 0
	// for x < 5 {
	// 	fmt.Println("x的值为：", x)
	// 	x++
	// }

	// for i :=0; i < 5; i++ {
	// 	fmt.Println("i的值为：", i)
	// }

	names := []string{"summer","lucy","henry","tom"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
  
	// for index, value := range names {
	// 	fmt.Printf("下标 %v 对应的值为 %v \n", index, value)
	// }

		for _, value := range names {
				fmt.Printf("遍历的值为 %v \n", value)
				value = "新名称"
			}
  // 不改变原始值
		fmt.Println(names)

}