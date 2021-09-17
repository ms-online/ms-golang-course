package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 获取输入值的方法
func getInput(prompt string ,r *bufio.Reader) (string,error){
  fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	// 阅读器
	reader := bufio.NewReader(os.Stdin)

	//输出询问
	// fmt.Print("创建账单名称：")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("创建账单名称：",reader)

	// 创建新账单
	 b := newBill(name)
	 fmt.Println("账单已创建：", b.name)

	 return b

}

//可选项
func promptOptions(b bill){
// 阅读器
	reader := bufio.NewReader(os.Stdin)

	opt,_ := getInput("可选项（a - 添加商品， s - 保存账单， t - 添加税额）",reader)
	fmt.Println(opt)
}

func main() { 
  mybill := createBill()
  promptOptions(mybill)

}
