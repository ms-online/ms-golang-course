package main

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
		 items: map[string]float64{},
		 tax: 0,
	 }
	 return b
 }
