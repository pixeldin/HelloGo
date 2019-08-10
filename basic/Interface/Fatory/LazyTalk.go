package Fatory

import (
	"errors"
	"fmt"
)

/*
	Duty接口简单实现, 可用作任意键机器人的使命插卡
	Hello(name string) string
	Talk(order string) (string, bool, error)
*/

//作为机器人使命的实现类, 必须实现所有接口(Hello和Talk)
type LazyDuty struct {
}

func (lt *LazyDuty) Hello(name string) string {
	return fmt.Sprintf("你好... %s ! -- 来自 lazy talk. zzzZZZ...", name)
}

func (lt *LazyDuty) Talk(order string) (saying string, end bool, err error) {
	switch order {
	case "":
		err := errors.New("指令未知")
		return "", false, err
	case "没有", "再见", "bye":
		saying = "再见！"
		end = true
		return
	default:
		saying = "emm... 我没听懂你说的, 请尝试下一个问题.  -- 来自 lazy talk. zzzZZZ..."
		return
	}
}
