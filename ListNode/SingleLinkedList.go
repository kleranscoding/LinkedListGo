package ListNode

import (
	"fmt"
	"strconv"
)

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *LinkedList) Print() {
	n, res := l.head, ""
	if n == nil {
		fmt.Println("empty linked list...")
		return
	}
	// read the head value and store in 'res'
	res = strconv.Itoa(n.val)
	n = n.next
	for n != nil {
		// add val to 'res'
		res += " -> " + strconv.Itoa(n.val)
		n = n.next
	}
	fmt.Println(res)
}

func (l *LinkedList) Insert(val int) {
	// TODO: insert a new node at tail

	// create a new node
	n := &Node {
		val: val,
	}
	// check if head is empty
	if (l.head == nil) {
		l.head = n
	} else {
		l.tail.next = n
	}
	// set tail to new node and increment size
	l.tail = n
	l.size++
}
