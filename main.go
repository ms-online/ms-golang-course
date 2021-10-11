package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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

	fmt.Println("-------------")
	fmt.Println("玩家的选择",playerChoice, "所对应的值为：", playerValue)
	fmt.Println("电脑的选择所对应的值为：",computerValue)



}