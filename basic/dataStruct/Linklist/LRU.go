package Linklist

/*
	LRU淘汰缓存, 实现思路:
	遍历原有链表, 与访问元素对比
	- 存在
	  将该元素移动至链表头部
	- 不存在
	  - 缓存空间足够
		新增该元素, 作为链表头部
	  - 缓存空间不足
		将该元素移动至链表头部, 淘汰链表尾部元素
 */
var LINK_SIZE = 3

func (l *LinkList) AddIntoLRU(ln *LinkNode) {
	node := l.GetNode(ln.value)
	if node != nil {
		l.DelNode(node)
	}
}

