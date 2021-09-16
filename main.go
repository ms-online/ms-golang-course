package main

import "fmt"

func updateName(x *string){
	*x = "lucy"
}

func updateFruit(y map[string]float64){
	y["peach"] = 2.88
}
	// 当变量作为参数传递给函数时，go 会生成一个“副本”

func main() { 
		// 类型A： string，int，boolean，float，array，struct
		// &:取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
		//* 号（前缀）来获取指针所指向的内容
	  name := "summer"

		// fmt.Println("name的内存地址为：", &name)
		
    // 手动添加指针
		m := &name
		// fmt.Println("name的内存地址为：", m)
		// fmt.Println("内存地址的值为：", *m)
	
		fmt.Println(name)
		updateName(m)
		fmt.Println(name)


		// 类型B：slice，map，function
		fruits := map[string]float64{
			"apple": 5.88 ,
			"orange": 4.68,
	}

	updateFruit(fruits)
	fmt.Println(fruits)
}


/*
|- - name - -|- - m - -|
|    0X001   |  0X002  |
|- - - - - - |- - - - -|
|  "summer"  | p0X001  |
|- - - - - - |- - - - -|
*/