package Linklist

import "testing"

func TestLinkList_TravelLinkList(t *testing.T) {
	ll := InitLinklist()
	l1 := &LinkNode{nil, 1}
	l2 := &LinkNode{nil, 2}
	l3 := &LinkNode{nil, 3}
	ll.Add(l1)
	ll.Add(l2)
	ll.Add(l3)

	ll.TravelLinkList()

	//ll.GetNode(2)

	ll.ReverseLinkList()
	ll.TravelLinkList()

}
