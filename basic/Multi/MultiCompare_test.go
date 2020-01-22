package main

import "testing"

var PIXEL_ARRAY []int

func init()  {
	for i := 1; i <= 100000; i++ {
		PIXEL_ARRAY = append(PIXEL_ARRAY, i)
	}
}
func BenchmarkSumWithSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumWithSingle(PIXEL_ARRAY)
	}
}

func BenchmarkSumWithMulti(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumWithMulti(PIXEL_ARRAY, 8)
	}
}
