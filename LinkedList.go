package main

import (
	"fmt"
	ll "LinkedList/ListNode"
)

func main() {
	l := ll.LinkedList{}
	l.Print()

	for i:= 1 ; i <= 20 ; i++ {
		l.Insert(i)
	}
	l.Print()
	searchVal := 9
	fmt.Println("search for", searchVal, "at index", l.Search(searchVal))
}