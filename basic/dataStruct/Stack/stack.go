package Stack

//栈的通用操作
type Stack interface {
	Top() interface{}
	Pop() interface{}
	Push(v interface{})
	IsEmpty() bool
	Clean()
}
