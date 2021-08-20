package stack

//Try to escape

// 不逃逸
func SpecifySizeAllocate()  {
	buf := make([]byte, 5)
	println(buf)
}

// 逃逸
func UnSpecifySizeAllocate()  {
	size := 5
	buf := make([]byte, size)
	println(buf)
}

func sliceMaker()  {
	size := 100
	s := make([]int, size)
	for idx, _ := range s {
			s[idx] = idx
	}
	print(s)
}

//go:noinline
func CreatePointer() *int  {
	//a := int(1)
	//return &a
	return new(int)
}
