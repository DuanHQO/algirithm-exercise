package Josephus

import "fmt"

type Person struct {
	num int
	code int
	next *Person
}

var head, tail *Person

func (ring *Person) AddJos(num, code int)  {
	current := &Person{code:code, num:num}
	if head == nil {
		head = current
		current.next = head
		tail = current
	} else {
		tail.next = current
		current.next = head
		tail = current
	}
}

func (ring *Person) RemoveJos(code int) int {
	if code == 1 {
		code = head.code
		fmt.Printf("%d  \n", head.num)
		head = head.next
		tail.next = head
		return code
	}

	for i := 1; i < code - 1; i++ {
		head = head.next
	}
	code = head.next.code
	fmt.Printf("%d ", head.next.num)
	tail = head
	head = head.next.next
	tail.next = head

	return code
}

func RunJos(ring *Person, m int) {
	code := m
	for head != tail {
		code = ring.RemoveJos(code)
	}
	fmt.Print(head.num)
}
