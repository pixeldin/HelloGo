package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	//Try to recover fatal err
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recover panic: ", p)
		}
	}()

	TypeDemo()

	fmt.Println("Call some wrong.")
	PanicCall()
	fmt.Println("Never come here.")
}

func TypeDemo() {
	sde := SelfDefineErr{"errMsg", 1}
	var sd, sdp interface{} = sde, &sde
	fmt.Println("SelfDefineErr's " + sd.(SelfDefineErr).errMsg)
	fmt.Println("SelfDefineErr pointer's " + sdp.(*SelfDefineErr).errMsg)

	//cannot assign to sd.(SelfDefineErr).errMsg
	//sd.(SelfDefineErr).errMsg = "2"
	t, ok := sd.(string)
	// false with nil value for assert type
	fmt.Println("## Type assert: ", ok, " + ", t)
	fmt.Println("## Type with reflect: ", reflect.TypeOf(sd))
	//equals to sd
	fmt.Println("## Value with reflect: ", reflect.ValueOf(sd))
}

func PanicCall() {
	somethingWrong()
}

func somethingWrong() {
	er := errors.New("Fatal error!")
	if er != nil {
		se := SelfDefineErr{er.Error(), 1}
		panic(se.Error())
	}
}

type SelfDefineErr struct {
	errMsg string
	tag    int
}

func (sde *SelfDefineErr) Error() string {
	return "Tag: " + fmt.Sprint(sde.tag) + ", msg:" + sde.errMsg
}
