package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func menu() {
	fmt.Println("1.积分抽奖")
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(9) + 1
	fmt.Println("	****************************")
	fmt.Println("	*	 积分抽奖系统	   *")
	fmt.Println("	****************************")
	fmt.Printf("	*   你抽到的积分为：%d分    *\n", r)
	fmt.Println("	*--------------------------*")
	fmt.Println("	* 请继续努力，按回车可退出 *")
	fmt.Println("	****************************")
	fmt.Println("2.任务抽取")
	fmt.Println("3.任务发布")
	fmt.Println("4.退出")
}
func main() {
	menu()

	bufio.NewReader(os.Stdin).ReadLine()
}
