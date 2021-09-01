package reflect

import (
	"io"
	"os"
	"reflect"
	"testing"
)

var (
	r io.Reader
	w io.Writer
	i interface{}
)

func TypeTransfer() {
	file, _ := os.OpenFile("", os.O_CREATE, 0)
	r = file

	// 编译失败
	// w = r
	// 类型断言
	var ok bool
	if w, ok = r.(io.Writer); !ok {
		panic("类型转换失败")
	}

	// 任意类型转interface{}无需断言
	i = ok

}

/*
	可设置(settability)有一点像地址可达(addressability), 这是一个反射对象可以修改实际创建该反射对象的值的属性,
	可设置与否取决于反射对象是否持有原始值（指针）.
*/
func TestRfc(t *testing.T) {
	sl1 := []int{1, 2, 3}
	sl2 := []int{1, 2, 3}
	// 不可寻址
	//t.Log(sl1==sl2)
	// 比较指针指向的底层数据是否相同
	t.Log(reflect.DeepEqual(sl1, sl2))

	mp1 := make(map[int]int, 0)
	mp2 := make(map[int]int, 0)
	//t.Log(mp1==mp2)
	t.Log(reflect.DeepEqual(mp1, mp2))

	fun1 := func() {
		print("hello")
	}
	fun2 := func() {
		print("hello")
	}
	t.Log(reflect.DeepEqual(fun1, fun2))

}
