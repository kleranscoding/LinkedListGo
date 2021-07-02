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

func (l *LinkedList) Search(val int) int {
	// search for val and return the index if not -1
	curr, index := l.head, 0
	if curr == nil {
		return -1
	}
	for curr != nil {
		if curr.val == val {
			return index
		}
		curr = curr.next
		index++
	}
	return -1
}
