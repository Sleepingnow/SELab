package main

// 结点数据结构
 type Node struct {
	next *Node	
	prev *Node
	data int
 }

// 双链表数据结构
type LinkTable struct {
	pHead *Node
	pTail *Node
	len int
}

// 创建链表
func CreatLinkTable() *LinkTable {
	var pLinkTable *LinkTable = new(LinkTable)
	if pLinkTable == nil {
		return nil
	} 
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.len = 0
	return pLinkTable
}

// 删除链表
func DetelLinkTable(pLinktable *LinkTable) int {
	// 删除空链表，返回-1
	if pLinktable == nil {
		return -1
	}
	pLinktable.pHead = nil
	pLinktable.pTail = nil
	pLinktable.len = 0
	return 0
}

// 尾部插入结点
func PushBack(pLinktable *LinkTable, pNode *Node) int {
	if pLinktable == nil || pNode == nil {
		return -1
	}
	if pLinktable.pHead == nil {
		pLinktable.pHead = pNode
		pLinktable.pTail = pNode
		pNode.next = nil
		pNode.prev = nil
	}else {
		pLinktable.pTail.next = pNode
		pNode.prev = pLinktable.pTail
		pLinktable.pTail = pNode
	}
	pLinktable.len++
	return 0
}

// 头部插入结点
func AddNode(pLinktable *LinkTable, pNode *Node) int {
	if pLinktable == nil || pNode == nil {
		return -1
	}
	if pLinktable.pHead == nil {
		pLinktable.pHead = pNode
		pLinktable.pTail = pNode
		pNode.next = nil
		pNode.prev = nil
	}else {
		pLinktable.pHead.prev = pNode
		pNode.next = pLinktable.pHead
		pLinktable.pHead = pNode
	}
	pLinktable.len++
	return 0
}

// 从双链表中删除指定结点
func DelNode(pLinktable *LinkTable, pNode *Node) int {
	if pLinktable == nil || pNode == nil {
		return -1
	}
	if pLinktable.pHead == pNode {
		pLinktable.pHead = pLinktable.pHead.next
		pLinktable.len--
		if pLinktable.pHead == nil {
			pLinktable.pTail = nil
		}
		return 0
	}
	var q *Node = pLinktable.pHead
	for q != nil {
		if q.next == pNode {
			q.next = q.next.next
			q = q.next.prev
			pLinktable.len--
			if pLinktable.len == 0 {
				pLinktable.pTail = nil
			}
			return 0
		}
		q = q.next
	}
	return -1
}

// 在双链表中查找指定结点
func SearchNode(pLinktable *LinkTable, pNode *Node) bool {
	if pLinktable == nil || pNode == nil {
		return false
	}
	var q *Node = pLinktable.pHead
	for q != nil {
		if q == pNode {
			return true
		}
		q = q.next
	}
	return false
}

// 获取双链表头结点
func GetLinkTableHead(pLinktable *LinkTable) *Node {
	if pLinktable == nil {
		return nil
	}
	return pLinktable.pHead
}

// 获取双链表尾结点
func GetLinkTableTail(pLinktable *LinkTable) *Node {
	if pLinktable == nil {
		return nil
	}
	return pLinktable.pTail
}

// 获取双链表长度
func GetLinkTableLen(pLinktable *LinkTable) int {
	if pLinktable == nil {
		return 0
	}
	return pLinktable.len
}