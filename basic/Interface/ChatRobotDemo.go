package main

import (
	ft "HelloGo/basic/Interface/Fatory"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type simpleRobot struct {
	name string
	duty ft.Duty
}

//构造方法
func NewSimpleRobot(name string, talkimp ft.Duty) ft.Chatbot {
	return &simpleRobot{
		name: name,
		duty: talkimp,
	}
}

func (robot *simpleRobot) Name() string {
	return robot.name
}

func (*simpleRobot) Begin() (string, error) {
	return "Please input your name", nil
}

func (*simpleRobot) End() error {
	return nil
}

func (*simpleRobot) ReportError(err error) string {
	return fmt.Sprintf("Err occurred: %s", err)
}

/*
	实现Talk接口
	type Talk interface {
		Hello(name string) string
		Talk(order string) (string, bool, error)
	}
 */
func (sr *simpleRobot) Hello(caller string) string {
	return fmt.Sprintf("hello " + caller + ", I'm " + sr.name + "\n" +
		"What't can I do for you?")
}

func (rb *simpleRobot) Talk(order string) (response string, end bool, err error) {
	od := strings.TrimSpace(order)
	// 假如肩负自己的使命, 则调用具体实现方法
	if rb.duty != nil {
		fmt.Print("Self implement: ")
		return rb.duty.Talk(od)
	}
	//默认使命
	switch order {
	case "":
		return
	case "没有", "bye":
		response = "bye！"
		end = true
		return
	default:
		response = "Default: sorry, nobody handle this, please retry later..."
		return
	}
}

func main() {
	var lt ft.LazyDuty
	//lt := new(Fatory.LazyDuty)
	//sr := &simpleRobot{
	//	"simpleRobot",
	//	&lt,
	//}
	//等同于
	sr := NewSimpleRobot("simpleRobot", &lt)

	//空实现, 调用默认方法
	//var t Talk
	//sr := simpleRobot{"sr", nil}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")

	readString, e := inputReader.ReadString('\n')
	if e != nil {
		fmt.Printf("An error occurred: %s\n", e)
		os.Exit(1)
	}
	caller := readString[:len(readString)-1]

	fmt.Println(sr.Hello(caller))


	for {

		s, e := inputReader.ReadString('\n')
		if e != nil {
			fmt.Printf("An error occurred: %s\n", e)
			os.Exit(1)
		}
		s = s[:len(s) - 1]
		s = strings.ToLower(s)

		response, end, _ := sr.Talk(s)
		fmt.Println(response)
		if end {
			break
		}

	}

}
