package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var PIXEL_ARRAY []int
	//初始化等差数列
	for i := 1; i <= 100000; i++ {
		PIXEL_ARRAY = append(PIXEL_ARRAY, i)
	}

	//单协程
	println(SumWithSingle(PIXEL_ARRAY))

	//多协程版本, 类似于Map-reduce
	println(SumWithMulti(PIXEL_ARRAY, 10))

}

func SumWithSingle(arr []int) int32 {
	var sum int32
	//for value := range arr {
	//	sum += value
	//	//print(value)
	//}
	for i := 0; i < len(arr); i++ {
		sum += int32(arr[i])
	}
	return sum
}

//每个协程均等计算自己部分的数列,
func SumWithMulti(arr []int, gNum int) int32 {
	var wg sync.WaitGroup
	wg.Add(gNum)

	var sum int32
	//注意切割长度需要向上取整,此处非侧重点, 使用整除切开原数组.
	//div := int(math.Ceil(float64(float64(len(arr)) / float64(gNum))))
	div := len(arr) / gNum
	for i := 0; i < gNum; i++ {
		Left := i * div
		Right := Left + div
		if i == gNum {
			Right = len(arr)
		}
		go func() {
			ps := 0
			for _, value := range arr[Left:Right] {
				ps += value
			}
			//由于仅有累加操作,可以用原子加实现互斥, 无需加锁.
			atomic.AddInt32(&sum, int32(ps))
			wg.Done()
		}()
	}

	//等待各个子协程计算完毕
	wg.Wait()
	return sum
}
