package main

import (
	ll "LinkedList/ListNode"
)

func main() {
	l := ll.LinkedList{}
	l.Print()

	for i:= 1 ; i <= 10 ; i++ {
		l.Insert(i)
	}
	l.Print()
}