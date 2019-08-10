package Fatory

import "errors"

type Duty interface {
	Hello(name string) string
	Talk(order string) (string, bool, error)
}

//机器人
type Chatbot interface {
	Name() string
	Begin() (string, error)
	//Talk 接口类型
	Duty
	ReportError(err error) string
	End() error
}

//机器人列表
var chatbot = map[string]Chatbot{}

//Common err
var (
	ErrInvalidChatBotName = errors.New("Invalid Chatbot Name")

	ErrInvalidChatBot = errors.New("Invalid Chatbot")

	ErrExistingChatBot = errors.New("Existing ChatBot")
)

//Register
func Register(cb Chatbot) error {
	if cb == nil {
		return ErrInvalidChatBot
	}
	name := cb.Name()
	if name == ""{
		return ErrInvalidChatBotName
	}
	//name already exist
	if _, ok := chatbot[name]; ok{
		return ErrExistingChatBot
	}
	chatbot[name] = cb
	return nil
}

func GetBot(name string) Chatbot {
	return chatbot[name]
}
