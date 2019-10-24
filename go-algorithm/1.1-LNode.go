package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

func Reverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var pre *LNode
	var cur *LNode
	next := node.Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

func RecursiveReverseChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := RecursiveReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}
func RecursiveReverse(node *LNode) {
	firstNode := node.Next
	newNode := RecursiveReverseChild(firstNode)
	node.Next = newNode
}

func InsertReverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var cur *LNode
	var next *LNode
	cur = node.Next.Next
	node.Next.Next = nil

	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}

func main() {
	head := &LNode{}
	fmt.Println("就地排序 通过改变指针指向来改变顺序")
	CreateNode(head, 8)
	PrintNode("逆序前 ", head)
	Reverse(head)
	PrintNode("逆序后 ", head)
	fmt.Println()

	head0 := &LNode{}
	fmt.Println("递归逆序")
	CreateNode(head0, 8)
	PrintNode("逆序前 ", head0)
	RecursiveReverse(head0)
	PrintNode("逆序后 ", head0)
	fmt.Println()

	head1 := &LNode{}
	fmt.Println("插入排序")
	CreateNode(head1, 8)
	PrintNode("排序前 ", head1)
	InsertReverse(head1)
	PrintNode("排序后 ", head1)
}

//就地排序
//逆序前 1 2 3 4 5 6 7
//逆序后 7 6 5 4 3 2 1

//递归逆序
//逆序前 1 2 3 4 5 6 7
//逆序后 7 6 5 4 3 2 1

//插入排序
//排序前 1 2 3 4 5 6 7
//排序后 7 6 5 4 3 2 1
