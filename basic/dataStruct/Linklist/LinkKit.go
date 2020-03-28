package Linklist

import (
	"fmt"
)

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

		//tmp := cur.next
		cur.next = node
		//node.next = tmp
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

func CopyNodeValue(l *LinkNode) *LinkNode {
	nn := new(LinkNode)
	nn.next = nil
	nn.value = l.value
	return nn
}

/*
	判断链表是否存在环
	快慢指针法, 指针一移动速度为1,指针二移动速度为2,
	当指针一追上指针二且非空时, 则链表存在环
	//应用:
	//如计算有序链表的中位数,
		- 当指针二刚好到达尾部, 则链表为奇数个, 指针一为中位数
		- 当指针二到达尾部前一个节点, 则链表为偶数个,
          指针一与下一节点求和均分为中位数
 */
func (l *LinkList)IsCircle() (cl bool) {
	cl = false
	if l.IsEmptyList() {
		return
	}
	p1 := l.head.next
	p2 := l.head.next.next
	for p2 != nil && p2.next != nil {
		if p1 == p2 {
			return true
		}
		//if p1.next == nil || p2.next == nil || p2.next.next == nil {
		//	return
		//}
		p1 = p1.next
		p2 = p2.next.next
	}
	return
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
		ln := CopyNodeValue(st.next)
		//ln := new(LinkNode)
		//ln.value = st.next.value

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
//方式2: 除去头结点, 从第二个节点开始遍历, 插入头结点后面
func (l *LinkList) ReverseLinkListLocal()  {
	if l.IsEmptyList() || l.length == 1 {
		return
	}
	cur := l.head.next
	for cur.next != nil {
		nxVal := CopyNodeValue(cur.next)
		cur.next = cur.next.next

		//insert after head
		nxVal.next = l.head.next
		l.head.next = nxVal
	}
}

//TODO:...方式3: 递归反转
func ReverseRecursive()  {
	//ReverseRecursive()
}


