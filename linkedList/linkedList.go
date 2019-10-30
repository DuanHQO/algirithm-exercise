package linkedList

import "fmt"

type Object interface {

}

type Node struct {
	Data int
	Next *Node
}

type List struct {
	headNode *Node
}

func (this *List) CopyFromArray(arr []int)  {
	for _, value := range arr {
		this.Append(value)
	}
}

func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	}else {
		return false
	}
}

func (this *List) Length() int {
	cur := this.headNode
	count := 0

	for cur != nil {
		count ++
		cur = cur.Next
	}
	return count
}

func (this *List) Add(data int) *Node  {
	node := &Node{Data:data}
	node.Next = this.headNode
	this.headNode = node
	return node
}

func (this *List) Append(data int)  {
	node := &Node{Data : data}
	if this.IsEmpty() {
		this.headNode = node
	}else {
		cur := this.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

func (this *List) Insert(index int, data int)  {
	if index <= 0 {
		this.Add(data)
	}else if index >= this.Length(){
		this.Append(data)
	}else{
		pre := this.headNode
		count := 0
		for count < (index - 1) {
			pre = pre.Next
			count ++
		}
		node := &Node{Data:data}
		node.Next = pre.Next
		pre.Next = node
	}
}

func (this *List) Remove(data int) {
	pre := this.headNode
	if pre.Data == data {
		this.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

func (this *List) RemoveAt(index int)  {
	pre := this.headNode
	if index <= 0{
		this.headNode = pre.Next
	} else if index >= this.Length() {
		fmt.Println("超出链表长度")
		return
	} else {
		count := 0
		for count != (index - 1) && pre.Next != nil {
			count ++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

func (this *List) GetHeadNode() *Node  {
	if !this.IsEmpty() && this.headNode != nil {
		return this.headNode
	}
	return nil
}

func (this *List) Contains(data Object) bool {
	cur := this.headNode
	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (this *List) ShowList()  {
	if !this.IsEmpty() {
		cur := this.headNode
		for {
			fmt.Printf("%d ", cur.Data)
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
		fmt.Println()
	}
}

func (this *List) RemoveAfter(node *Node)  {
	if node == nil {
		return
	}
	if &node.Next == nil {
		return
	}
	node.Next = nil
}

func (this *List) Reverse()  {
	if this.IsEmpty() {
		return
	}
	first := this.headNode
	var previous *Node
	for {
		if first != nil {
			second := first.Next
			first.Next = previous
			previous = first
			first = second
		}
	}
}

func (this *List) MergeSort()  {
	this.headNode = mergeSort(this.headNode)
	this.ShowList()
}

func mergeSort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMidNode(head)
	second := mid.Next
	mid.Next = nil

	left := mergeSort(head)
	right := mergeSort(second)

	return merge(left, right)
}

func merge(left, right *Node) *Node {

	dummy := new(Node)
	tail := dummy

	for left != nil && right != nil {
		 if left.Data < right.Data {
		 	tail.Next = left
		 	left = left.Next
		 } else {
		 	tail.Next = right
		 	right = right.Next
		 }
		 tail = tail.Next
	}

	if left == nil {
		tail.Next = right
	} else {
		tail.Next = left
	}

	return dummy.Next
}

func getMidNode(head *Node) *Node  {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	faster := slow.Next
	for faster != nil && faster.Next != nil {
		slow = slow.Next
		faster = faster.Next.Next
	}
	return slow
}