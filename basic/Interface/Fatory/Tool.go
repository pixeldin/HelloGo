package Fatory

import (
	"errors"
	"fmt"
)

//通用使命: SayHello, Talk
type Duty interface {
	Hello(name string) string
	Talk(order string) (string, bool, error)
}

//机器人
type Chatbot interface {
	Name() string
	Begin() (string, error)
	//机器人的使命插槽, Duty 接口类型 (可由具体Duty接口实现来替换)
	Duty
	ReportError(err error)
	End() error
}

//机器人列表
var chatbot = map[string]Chatbot{}

//Common err
var (
	ErrInvalidChatBotName = errors.New("Invalid Chatbot Name")
	ErrInvalidChatBot     = errors.New("Invalid Chatbot")
	ErrExistingChatBot    = errors.New("Existing ChatBot")
)

//Register
func Register(cb Chatbot) error {
	if cb == nil {
		return ErrInvalidChatBot
	}
	name := cb.Name()
	if name == "" {
		return ErrInvalidChatBotName
	}
	//name already exist
	if _, ok := chatbot[name]; ok {
		return ErrExistingChatBot
	}
	chatbot[name] = cb
	return nil
}

func GetBot(name string) Chatbot {
	robot := chatbot[name]
	if robot == nil {
		fmt.Println("Sorry, can't find your robot." +
			" We call a default implement robot for you.")
		return NewSimpleRobot("default", nil)
	}
	fmt.Println("Your robot is coming, ROBOT " + name)
	return robot
}
