package stack

//Try to escape

//go:noinline
func CreatePointer() *int  {
	//a := int(1)
	//return &a
	return new(int)
}
