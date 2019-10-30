package Steque

import "algorithm/linkedList"

type Queue struct {
	linkedList.List
}

func (this *Queue) Push(data int)  {
	this.Append(data)
}

func (this *Queue) Pop() *linkedList.Node {
	if this.GetHeadNode() == nil {
		panic("this queue is nil")
	}
	node := this.GetHeadNode()
	this.RemoveAt(0)
	return node
}

func (this *Queue) Enqueue() {

}
