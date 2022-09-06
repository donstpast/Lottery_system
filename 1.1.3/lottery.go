package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//1.1.3需求：。完成任务抽取功能初步，即可以进去之后随机抽取一个内容
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
	var operate string
	var tasks = [...]string{"考研数学复习一小时", "考研政治复习一小时", "考研英语刷题一小时", "考研计算机复习一小时", "考研计算机刷题一小时", "阅读30分钟", "改错题半小时", "写代码半小时", "看技术书/文章半小时", "改错题半小时"}
	rand.Seed(time.Now().UnixNano())
	var num int = len(tasks)
	thing := rand.Intn(num) + 1
	fmt.Println("	********************************************************")
	fmt.Println("	*	                 任务抽取    	               *")
	fmt.Println("	********************************************************")
	fmt.Printf("	                 你抽到的任务为 ：%v             \n ", tasks[thing])
	fmt.Println("	*------------------------------------------------------*")
	fmt.Println("	*        	     请认真完成任务     	       *")
	fmt.Println("	*------------------------------------------------------*")
	fmt.Println("	*            输入r可以重新抽取，输入q可以返回主菜单    *")
	fmt.Println("	*------------------------------------------------------*")
	fmt.Println("	********************************************************")
	for {
		fmt.Println("请输入您的操作：")
		fmt.Scanln(&operate)
		if operate == "q" {
			main()
		} else if operate == "r" {
			task()
		} else {
			fmt.Println("您的输入有误，请重新输入！！！")
		}
	}
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
