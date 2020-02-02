package main

import (
	"fmt"
	"reflect"
)

type TypeError struct {
	//ty string
	//上下文
	context string
	Type    reflect.Type
	trace   string
}

func (tye *TypeError) Error() string {
	return fmt.Sprintf("Type of TypeError %s, contex: %s, trace[%s]",
		tye.Type, tye.context, tye.trace)
}

type SizeError struct {
	//size int
	//context
	context string
	Type    reflect.Type
}

func (sie *SizeError) Error() string {
	return fmt.Sprintf("SizeError context: %s", sie.context)
}

type UserError struct {
	context string
	Type    reflect.Type
}

func (tie *UserError) Error() string {
	return fmt.Sprintf("UserError contex: %s", tie.context)
}
