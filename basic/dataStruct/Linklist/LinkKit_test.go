package Linklist

import (
	"fmt"
	"strings"
	"testing"
)

func TestLinkList_TravelLinkList(t *testing.T) {
	ll := InitLinklist()
	l1 := &LinkNode{nil, 1}
	l2 := &LinkNode{nil, 2}
	l3 := &LinkNode{nil, 3}
	l4 := &LinkNode{nil, 4}
	ll.Add(l1)
	ll.Add(l2)
	ll.Add(l3)
	ll.Add(l4)

	ll.TravelLinkList()

	//ll.GetNode(2)

	//ll = ll.ReverseLinkListByNewLink()
	//ll.TravelLinkList()
	ll.ReverseLinkListLocal()
	ll.TravelLinkList()

}

func TestStringsRepeat(t *testing.T)  {
	repeat := strings.Repeat("/:param", 2)
	fmt.Println(repeat)
	const maxParamCount uint8 = ^uint8(0)

	fmt.Println(maxParamCount)
}