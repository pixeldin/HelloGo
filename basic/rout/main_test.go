package main

import "testing"

func TestGetGID(t *testing.T) {
	//GetGID()
	//getG()
}

// BenchmarkGId-8   	1000000000	         0.0005081 ns/op
func BenchmarkGRtId(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// runtime获取协程id
		GetGID()
	}
}

// BenchmarkLog-8   	1000000000	         0.05731 ns/op
func BenchmarkGoId(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// 汇编方式获取
		GoId()
	}
}
