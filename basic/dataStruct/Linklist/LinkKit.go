package Linklist

import "fmt"

/*
	单链表操作
*/

//链表节点, 下一个成员地址/当前value
type LinkNode struct {
	next  *LinkNode
	value interface{}
}

//链表, 队头节点/长度
type LinkList struct {
	head   *LinkNode
	length int
}

func InitLinklist() *LinkList {
	ln := new(LinkNode)
	ln.next = nil
	ln.value = nil
	ll := &LinkList{ln, 1}
	return ll
}

func (l *LinkList) IsEmptyList() bool {
	return l.length == 1
}

func (l *LinkList) Add(node *LinkNode) {
	if l.IsEmptyList() {
		l.head.next = node
	} else {
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}

		tmp := cur.next
		cur.next = node
		node.next = tmp
	}
	l.length++
}

//在节点前添加
func (l *LinkList) AddBefore(node *LinkNode, v interface{}) {

}

func (l *LinkList) AddAfter(node *LinkNode, v interface{}) {

}

//删除
func (l *LinkList) DelNode(node *LinkNode) {
	if l.IsEmptyList() {
		return
	}
	cur := l.head
	for cur.next != nil {
		if cur.next == node {
			cur.next = node.next
			l.length--
			//free(node)
			return
		}
		cur = cur.next
	}
}

func (l *LinkList) TravelLinkList() {
	var vs string = "|Head->"
	for p := l.head; p.next != nil; {
		p = p.next
		vs += fmt.Sprintf("%v->", p.value)
	}
	vs += "Tail|"
	fmt.Printf("LinkList: %s\n", vs)
}

func (l *LinkList) GetNode(value interface{}) *LinkNode {
	if l.length > 1 {
		for cur := l.head; cur.next != nil; {
			cur = cur.next
			if cur.value == value {
				return cur
			}
		}
	}
	return nil
}

/*
	链表反转
*/
// 方式1: 遍历原链表, 在新创建链表头结点之后逐个插入.
func (l *LinkList) ReverseLinkListByNewLink() (nl *LinkList){
	if l.IsEmptyList() {
		return
	}
	nl = InitLinklist()
	// head/st ->  -> next
	st := l.head
	for st.next != nil {
		ln := new(LinkNode)
		ln.value = st.next.value

		old := nl.head.next
		//if old != nil {
		ln.next = old
		//}
		nl.head.next = ln
		nl.length++

		st = st.next
	}
	return nl
}
//方式2: 就地反转
func (l *LinkList) ReverseLinkList()  {

}
