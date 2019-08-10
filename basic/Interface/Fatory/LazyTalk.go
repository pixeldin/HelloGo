package Fatory

import "fmt"

/*
	Talk 接口简单实现
	Hello(name string) string
	Talk(order string) (string, bool, error)
*/

//作为机器人使命的实现类, 必须实现所有接口(Hello和Talk)
type LazyDuty struct {
}

func (lt *LazyDuty) Hello(name string) string {
	return fmt.Sprintf("你好, %s ! -- 来自 lazy talk.", name)
}

func (lt *LazyDuty) Talk(order string) (saying string, end bool, err error) {
	switch order {
	case "":
		return
	case "没有", "再见", "bye":
		saying = "再见！"
		end = true
		return
	default:
		saying = "对不起，我没听懂你说的, 换个指令?"
		return
	}
}
