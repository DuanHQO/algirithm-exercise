package MoveToFront

import (
	"fmt"
	"io"
)

type Node struct {
	char byte
	next *Node
}

type MoveStack struct {
	head *Node
}

func (this *MoveStack) Push(char byte) {
	node := &Node{char:char}
	if this.Judge(char) {
		err := this.Remove(char)
		if err != nil {
			panic("remove char failed")
		}
	}
	if this.head != nil {
		node.next = this.head
	}
	this.head = node
}

func (this *MoveStack) Judge(char byte) bool {

	cur := this.head
	for cur != nil {
		if cur.char == char {
			return true
		} else {
			cur = cur.next
		}
	}

	return false
}

func (this *MoveStack) Remove(char byte) error {
	cur := this.head
	if cur.char == char {
		this.head = cur.next
		return nil
	}
	for cur.next != nil {
		if cur.next.char == char {
			cur.next = cur.next.next
			return nil
		} else {
			cur = cur.next
		}
	}
	return io.EOF
}

func (this *MoveStack) ShowList() {
	cur := this.head
	for cur != nil {
		fmt.Printf("%c ", cur.char)
		cur = cur.next
	}
}