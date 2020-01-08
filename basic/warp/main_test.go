package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestWeekNum(t *testing.T) {
	//pt := time.Now()
	pt1 := time.Date(2019, 12, 29, 13, 0, 0, 0, time.UTC)
	pt2 := time.Date(2019, 12, 30, 13, 0, 0, 0, time.UTC)
	_, weekNum1 := pt1.ISOWeek()
	_, weekNum2 := pt2.ISOWeek()
	//month := pt.Month()
	fmt.Println("P1 week of year: " + fmt.Sprint(weekNum1))
	fmt.Println("P2 week of year: " + fmt.Sprint(weekNum2))
	//fmt.Println("Week of year: " + fmt.Sprint(month))
	////周一中午开赛前数据仍然视为上一周数据
	//if pt.Weekday() == 1 && pt.Hour() < 12 {
	//
	//}
}

func TestMonth(t *testing.T) {
	pt1 := time.Date(2019, 1, 29, 13, 0, 0, 0, time.UTC)
	lastMonth := pt1.Month() - 1
	if lastMonth == 0 {
		lastMonth = 12
	}
	fmt.Println("lastMonth",  int(lastMonth))
}

func TestSome(t *testing.T) {
	//i := 1
	//fmt.Printf("Test i add: %x\n", &i)
	//fmt.Printf("Test i add: %p\n", &i)
	//fmt.Println("Test i add: ", &i)
	//var j *int
	//if j != nil {
	//	fmt.Println(*j)
	//}

	var itn interface{}
	//var itn *Person
	//itf = GetSomeNotNil()
	itn = GetSomeWithWrapped()
	//itn = nil

	//interface 类型如果类型不为interface, 那么则不为空
	if itn == nil {
		fmt.Println("Nil interface... type:", reflect.TypeOf(itn))
	} else {
		fmt.Println("Type:", reflect.TypeOf(itn))
	}

	s := []int{1, 2, 3}
	for i, v := range s {
		if i == 1 {
			s[0] = 6
			fmt.Println(s)
		}
		fmt.Println(i, v)
	}
	fmt.Println("After travel :", s)
}

func BenchmarkWithNoBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("Abc...")
		//time.Sleep(2000)
	}
}
