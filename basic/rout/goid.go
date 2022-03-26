package main

import (
	"reflect"
	"runtime"
	"strings"
)

// 记录各个版本的偏移量
var offsetDictMap = map[string]int64{
	"go1.12":   152,
	"go1.12.1": 152,
	"go1.12.2": 152,
	"go1.12.3": 152,
	"go1.12.4": 152,
	"go1.12.5": 152,
	"go1.12.6": 152,
	"go1.12.7": 152,
	"go1.13":   152,
	"go1.14":   152,
	"go1.16.9": 152,
}

// offset for go1.12
var goid_offset uintptr = 152

//go:nosplit
func getG() interface{}

func init() {
	version := runtime.Version()
	g := getG()
	if nil != g {
		g_type := reflect.TypeOf(g)
		goid_field, ok := g_type.FieldByName("goid")
		if !ok {
			panic("goid not found. version:" + version)
		}
		goid_offset = goid_field.Offset
	} else {
		ok := false
		for k, v := range offsetDictMap {
			if version == k || strings.HasPrefix(version, k) {
				goid_offset = uintptr(v)
				ok = true
				break
			}
		}
		if !ok {
			panic("goid not found. version:" + version)
		}
	}
}

func GoId() int64
