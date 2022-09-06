package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//今日完善需求：积分抽取完可以返回上一级，并可以重新抽取。
var choice int8

func menu() {
	fmt.Println("1.积分抽奖")
	fmt.Println("2.任务抽取")
	fmt.Println("3.退出")
}
func points() {
	var operate string
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(9) + 1
	fmt.Println("	****************************")
	fmt.Println("	*	 积分抽奖系统	   *")
	fmt.Println("	****************************")
	fmt.Printf("	*   你抽到的积分为：%d分    *\n", r)
	fmt.Println("	*--------------------------*")
	fmt.Println("	*        请继续努力!!!     *")
	fmt.Println("	*--------------------------*")
	fmt.Println("	*     输入r可以重新抽取    *")
	fmt.Println("	*--------------------------*")
	fmt.Println("	*    输入q可以返回主菜单   *")
	fmt.Println("	*--------------------------*")
	fmt.Println("	****************************")
	for {
		fmt.Println("请输入您的操作：")
		fmt.Scanln(&operate)
		if operate == "q" {
			main()
		} else if operate == "r" {
			points()
		} else {
			fmt.Println("您的输入有误，请重新输入！！！")
		}
	}

}
func task() {
	fmt.Println("2.任务抽取")
}

func main() {
	menu()
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		points()
	case 2:
		task()
	case 3:
		os.Exit(1)
	default:
		fmt.Println("您的选择有误，请重新选择")
	}

	bufio.NewReader(os.Stdin).ReadLine()
}
