package main

import (
	ft "HelloGo/basic/Interface/Fatory"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	GET_ROBOT_COMMAND = "Please input your robot name:"
	GET_USER_NAME     = "Please input your name:"
	GET_USER_COMMAND  = "What't can I do for you?"
)

//获取用户指令
func GetManInput(hint string) (input string) {
	fmt.Println(hint)
	inputReader := bufio.NewReader(os.Stdin)
	input, e := inputReader.ReadString('\n')
	if e != nil {
		fmt.Printf("An error occurred: %input\n", e)
		os.Exit(1)
	}
	input = input[:len(input)-1]
	return
}

func main() {
	//使命具体实现
	//返回接口引用 lt := new(ft.LazyDuty)
	var lt ft.LazyDuty
	//使命芯片接入机器人
	sr := ft.NewSimpleRobot("lazy", &lt)

	//注册到机器人工具
	ft.Register(sr)

	//空实现, 调用默认方法
	//var t Talk
	//sr := simpleRobot{"sr", nil}

	//呼叫机器人
	robotName := GetManInput(GET_ROBOT_COMMAND)
	robot := ft.GetBot(robotName)

	//获取用户名称
	userName := GetManInput(GET_USER_NAME)
	fmt.Println(robot.Hello(userName))

	for {

		//获取用户指令
		s := GetManInput(GET_USER_COMMAND)
		order := strings.ToLower(s)

		response, end, err := robot.Talk(order)
		if err != nil {
			robot.ReportError(err)
		}
		fmt.Println(response)
		if end {
			break
		}

	}

}
