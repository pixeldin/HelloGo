package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"
)

var (
	TYE = "typeErr"
	SE  = "sizeErr"
	UE  = "userErr"
)

func main() {
	typeErr := CreateWithDiffError(TYE)
	seErr := CreateWithDiffError(SE)
	userErr := CreateWithDiffError(UE)

	ParseErr(typeErr)
	ParseErr(seErr)
	ParseErr(userErr)

	wrapErr := WrapErr(userErr)
	//fmt.Println(wrapErr)
	ParseErr(wrapErr)
}

/*
	模拟不同场景产生不同错误
*/
func CreateWithDiffError(instruction string) error {
	switch instruction {
	case "typeErr":
		//上下文添加备注
		return &TypeError{"Lack of energy.", reflect.TypeOf(TypeError{}), getStackTrace()}
	case "sizeErr":
		//上下文添加时间信息
		return &SizeError{"time:" + time.UnixDate, reflect.TypeOf(SizeError{})}
	case "userErr":
		//上下文添加用户
		return &UserError{"UserErr with selfContext: @pixelpig.",
			reflect.TypeOf(UserError{})}
	default:
		return errors.New("UnknownError")
	}
}

func ParseErr(err error) {
	if err != nil {
		switch e := err.(type) {
		case *TypeError:
			log.Printf("ErrType[%v] Context[%s]\n, Trace[%s]\n", e.Type, e.context, e.trace)
			//TODO: Handle 类型错误
			break
		case *UserError:
			log.Printf("ErrType[%v] Context[%s]\n", e.Type, e.context)
			//TODO: Handle 用户错误
			break
		case *SizeError:
			log.Printf("ErrType[%v] Context[%s]\n", e.Type, e.context)
			//TODO: Handle 容量错误
			break
		default:
			log.Println(err)
		}
	}
}

/*
	https://www.komu.engineer/blogs/golang-stacktrace/golang-stacktrace
	获取当前执行点的栈信息
*/
//Package errors provides ability to annotate you regular Go errors with stack traces.
func getStackTrace() string {
	stackBuf := make([]uintptr, 50)
	length := runtime.Callers(3, stackBuf[:])
	stack := stackBuf[:length]

	trace := ""
	frames := runtime.CallersFrames(stack)
	for {
		frame, more := frames.Next()
		trace = trace + fmt.Sprintf("\n\tFile: %s, Line: %d. Function: %s",
			frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
	return trace
}

/*
	包装错误
*/
func WrapErr(err error) error {
	return fmt.Errorf("Wrap with shell: %v", err)
}
