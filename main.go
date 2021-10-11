package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// 创建常量
const (
	ROCK = 0
	PAPER = 1
	SCISSORS = 2
)
func main() {
	// 设置随机数种子,保证每次随机都是随机的
	rand.Seed(time.Now().UnixNano())
  playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(2)

	// 读取标准输入
	reader := bufio.NewReader(os.Stdin)

	//清空屏幕
	clearScreen()

	// 游戏提示
	fmt.Print("请输入rock,paper,scissors -> ")
  
	// 读取用户输入的选项
	playerChoice, _ = reader.ReadString('\n')
	// 清理用户输入
	playerChoice = strings.Replace(playerChoice,"\n", "", -1)

	// 保存用户选项所对应的值
	if playerChoice == "rock"{
		playerValue = ROCK
	}else if playerChoice == "paper"{
			playerValue = PAPER
	}else if playerChoice == "scissors"{
			playerValue = SCISSORS
	}

	fmt.Println("玩家的选择",playerChoice)

	// 输出电脑随机生成的选项
	switch computerValue {
	  case ROCK :
		  fmt.Println("电脑的选择为ROCK")
		case PAPER :
			fmt.Println("电脑的选择为PAPER")
		case SCISSORS :
			fmt.Println("电脑的选择为SCISSORS")
	}

	fmt.Println("-------------")
	
 // 本回合结论
  if playerValue == computerValue {
		fmt.Println("本回合为平局！")
	}else {
		switch playerValue {
		case ROCK :
			if computerValue == PAPER {
				fmt.Println("本回合电脑获胜！")
			}else {
				fmt.Println("本回合玩家获胜！")
			}
		case PAPER :
			if computerValue == SCISSORS {
				fmt.Println("本回合电脑获胜！")
			}else {
				fmt.Println("本回合玩家获胜！")
			}
		case SCISSORS :
			if computerValue == ROCK {
				fmt.Println("本回合电脑获胜！")
			}else {
				fmt.Println("本回合玩家获胜！")
			}
		default:
			fmt.Println("无效输入")
		}
	}


}


// 清空屏幕
func clearScreen(){
	// runtime提供go运行时环境的操作
	//GOOS判断电脑的操作系统
	if strings.Contains(runtime.GOOS,"windows"){
	cmd := 	exec.Command("cmd","/c","cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	} else {
	cmd := 	exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	}
}