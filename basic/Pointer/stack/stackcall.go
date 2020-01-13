package stack

const size = 2048

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

func SomeExtrem() {
	a := new(int)
	ap := &a
	x := &ap
	//fmt.Printf("a: %p, ap: %p, x: %p\n", a, ap, x)
	println("a: ", a, " ap: ", ap, " x:", x)
}
