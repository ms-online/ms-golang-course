package main

import "fmt"

// 自定义数据类型
type bill struct {
	name string
	items map[string]float64
	tax float64
}

// 创建新账单
 func newBill(name string) bill{
	 b := bill{
		 name: name,
		 items: map[string]float64{"apple":5.88,"orange": 2.56},
		 tax: 0,
	 }
	 return b
 }

 // 格式化账单
 func (b bill) format() string{
   fs := "Bill Details: \n"
	 var total float64 = 0

	 // 遍历商品集合
    for k, v := range b.items{
			fs += fmt.Sprintf("%-25v ...$%v \n", k + ":", v)
			total += v
		}

		//总金额
		fs += fmt.Sprintf("%-25v ...$%0.2f","total:", total)

		return fs

 }
