package main

const size = 2048

func main() {
	//var v int = 1
	vp := new(int)
	*vp = 1
	//fmt.Printf("# Main frame: Value of v:\t\t %v, address: %p\n", v, &v)
	//println("# Main frame: Value of v:\t\t", v, " address: ", &v)
	println("# Main frame: Value of v:\t\t", *vp, " address: ", vp)

	//PassValue(v, &v)
	PassValue(*vp, vp)

	//fmt.Printf("# Main frame: Value of v:\t\t %v, address: %p\n", v, &v)
	//println("# Main frame: Value of v:\t\t", v, " address: ", &v)
	println("# Main frame: Value of v:\t\t", *vp, " address: ", vp)


	//SomeExtrem()

	//fmt.Println("==================")
	//s := "HELLO"
	//stackCopy(0, &s, [size]int{})
}

// stackCopy recursively runs increasing the size
// of the stack.
func stackCopy(c int, s *string, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(c, s, a)
}

func PassValue(fv int, addV *int) {
	// fv 的地址只属于该函数, 由该函数栈分配
	//fmt.Printf("# Func frame: Value of fv:\t\t %v, address: %p\n", fv, &fv)
	println("# Func frame: Value of fv:\t\t ", fv, " address: ", &fv)
	//本次修改只在该函数生效
	fv = 0
	//fmt.Printf("# Func frame: Value of fv:\t\t %v, address: %p\n", fv, &fv)
	println("# Func frame: Value of fv:\t\t ", fv, " address: ", &fv)

	/*
	 *	根据main函数传入的全局地址, 对指针执行操作外部是可见的,
	 *  因为改指针操作的都是同一个内存块的内容
	 */
	*addV++
	//fmt.Printf("# Func frame: Value of addV:\t %v, address: %p\n", *addV, addV)
	println("# Func frame: Value of addV:\t ", *addV, " address: ", addV)
}

func SomeExtrem() {
	a := new(int)
	ap := &a
	x := &ap
	//fmt.Printf("a: %p, ap: %p, x: %p\n", a, ap, x)
	println("a: ", a, " ap: ", ap, " x:", x)
}
