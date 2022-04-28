package list

import (
	"fmt"
	"strings"
)

type List struct {
	root *Elem
}

type Elem struct {
	Val        int
	next, prev *Elem
}

func New() *List {
	var l List

	l.root = &Elem{}
	l.root.next = l.root
	l.root.prev = l.root

	return &l
}

func (l *List) Push(e Elem) *Elem {
	e.prev = l.root
	e.next = l.root.next

	l.root.next.prev = &e
	l.root.next = &e

	return &e
}

func (l *List) String() string {
	s := make([]string, 0)

	for e := l.root.next; e != l.root; e = e.next {
		s = append(s, fmt.Sprintf("%d", e.Val))
	}

	return strings.Join(s, " ")
}

func (l *List) Pop() *List {
	l.root.prev, l.root.next = l.root.next, l.root.next.next
	return l
}

func (l *List) Reverse() *List {
	var prev, current *Elem = nil, l.root

	for current != nil {
		next := current.next
		current.next, current.prev = prev, next
		prev, current = current, next
	}

	return l
}
