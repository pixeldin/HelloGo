package Fatory

import (
	"errors"
	"fmt"
	"strings"
)

type simpleRobot struct {
	name string
	duty Duty
}

//构造方法
func NewSimpleRobot(name string, dutyImpl Duty) Chatbot {
	return &simpleRobot{
		name: name,
		duty: dutyImpl,
	}
}

func (robot *simpleRobot) Name() string {
	return robot.name
}

func (*simpleRobot) Begin() (string, error) {
	return "", nil
}

func (*simpleRobot) End() error {
	return nil
}

func (sr *simpleRobot) ReportError(err error) {
	errString := fmt.Sprintf("Err occurred from: %s, err: %s", sr.name, err)
	fmt.Println(errString)
}

/*
	实现Duty默认接口
	type Duty interface {
		Hello(name string) string
		Talk(order string) (string, bool, error)
	}
*/
func (sr *simpleRobot) Hello(caller string) string {
	// 假如肩负了具体使命, 则调用具体实现方法
	if sr.duty != nil {
		return sr.duty.Hello(caller)
	}
	//默认使命
	return fmt.Sprintf("Hello " + caller + ", I'm " + sr.name)
}

func (rb *simpleRobot) Talk(order string) (response string, end bool, err error) {
	od := strings.TrimSpace(order)
	// 假如肩负了具体使命, 则调用具体实现方法
	if rb.duty != nil {
		fmt.Print("Self implement: ")
		return rb.duty.Talk(od)
	}
	//默认使命
	switch order {
	case "":
		err := errors.New("Unknown command")
		return "", false, err
	case "没有", "bye":
		response = "bye！"
		end = true
		return
	default:
		response = "Default: sorry, nobody handle this, " +
			"please retry another command..."
		return
	}
}
