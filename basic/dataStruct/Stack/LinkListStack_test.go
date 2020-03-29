package Stack

import (
	"fmt"
	"testing"
)

func TestLinkStack(test *testing.T)  {
	ls := InitStack()
	fmt.Printf("Empty stack? %v\n", ls.IsEmpty())

	ls.Push(1)
	ls.Push(2)
	ls.Push(3)

	fmt.Printf("After push, still empty stack? %v\n", ls.IsEmpty())
	fmt.Printf("Top value: %v\n", ls.Top())

	ls.Travel()
	fmt.Printf("Pop value: %v,\nafter pop,", ls.Pop())
	ls.Travel()
}
