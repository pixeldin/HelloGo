package stack

//Try to escape

func SpecifySizeAllocate()  {
	buf := make([]byte, 5)
	println(buf)
}

func UnSpecifySizeAllocate(size int)  {
	buf := make([]byte, size)
	println(buf)
}

//go:noinline
func CreatePointer() *int  {
	//a := int(1)
	//return &a
	return new(int)
}
