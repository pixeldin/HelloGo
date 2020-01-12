package main

import "fmt"

const size = 2048

func main() {
	//var v int = 1
	////var pointer = new(int)
	//fmt.Printf("#Main frame: Value of v: %v, address: %p\n", v, &v)
	////fmt.Printf("#Main frame: Value of pointer address: %p\n", pointer)
	////*pointer = 1
	//PassSomethingIntoHere(v, &v)
	//SomeExtrem()

	fmt.Println("==================")
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

// stackCopy recursively runs increasing the size
// of the stack.
func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(s, c, a)
}

func PassSomethingIntoHere(a int, pointA *int) {
	fmt.Printf("#Func frame: Value of param: %v, address: %p\n", a, &a)
	//a++
	//fmt.Printf("#Func frame2: Value of param: %v, address: %p\n", a, &a)
	//*pointA = 1
	fmt.Printf("#Func frame: Value of pointer address: %p\n", pointA)
}

func SomeExtrem() {
	a := new(int)
	ap := &a
	x := &ap
	fmt.Printf("a: %p, ap: %p, x: %p\n", a, ap, x)
}
