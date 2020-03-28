package Linklist

import (
	"fmt"
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

	//l4.next = l2
	cl := ll.IsCircle()
	if cl {
		fmt.Println("是个环链表.")
	} else {
		fmt.Println("不是环链表.")
	}
	//ll.Add(l4)

	fmt.Println("=========\n原链表: ")
	ll.TravelLinkList()


	//ll.GetNode(2)

	fmt.Println("反转1: ")
	ll = ll.ReverseLinkListByNewLink()
	ll.TravelLinkList()

	fmt.Println("反转2: ")
	ll.ReverseLinkListLocal()
	ll.TravelLinkList()

}

