package list

import "math"

type Node[V any] struct {
	Value      V
	next, prev *Node[V]
}

type List[V any] struct {
	head, tail *Node[V]
	len        int
}

func (n *Node[V]) Next() *Node[V] {
	return n.next
}

func (n *Node[V]) Prev() *Node[V] {
	return n.prev
}

//New creates new List with passed args.
func New[V any](args ...V) *List[V] {
	newList := new(List[V])
	for _, v := range args {
		newList.Add(v)
	}
	return newList
}

func (l *List[V]) Len() int {
	return l.len
}

//Add adds new item to collection.
//Returns pointer to new node with added value.
func (l *List[V]) Add(value V) *Node[V] {
	if l.len == 0 {
		l.head = &Node[V]{value, nil, nil}
		l.tail = l.head
	} else {
		l.tail.next = &Node[V]{value, nil, l.tail}
		l.tail = l.tail.next
	}
	l.len++
	return l.tail
}

//Remove removes node from collection.
//Connects deleted node's neighbours.
func (l *List[V]) Remove(n *Node[V]) {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		l.head = l.head.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		l.tail = l.tail.prev
	}

	n.next = nil
	n.prev = nil
	l.len--
}

//Reduce calls cb for every collection's item.
//Returns accumulated value.
func (l *List[V]) Reduce(cb func(acc, el V) V) V {
	var res V
	for n := l.head; n != nil; n = n.next {
		res = cb(res, n.Value)
	}
	return res
}

//ForEach calls cb for every collection's item.
//Modifies current list.
func (l *List[V]) ForEach(cb func(el *V, i int)) {
	i := 0
	for n := l.head; n != nil; n = n.next {
		cb(&n.Value, i)
		i++
	}
}

//Map calls cb for every collection's item.
//Returns new list.
func (l *List[V]) Map(cb func(el V, i int) V) *List[V] {
	i := 0
	newList := New[V]()
	for n := l.head; n != nil; n = n.next {
		newList.Add(cb(n.Value, i))
		i++
	}

	return newList
}

//Head returns pointer to the first node.
func (l *List[V]) Head() *Node[V] {
	return l.head
}

//Tail returns pointer to the last node.
func (l *List[V]) Tail() *Node[V] {
	return l.tail
}

//At returns node at i position.
//Negative index means searching item from the end.
func (l *List[V]) At(i int) *Node[V] {
	if i < 0 {
		i = int(math.Abs(float64(l.len + i)))
	}
	i = i % l.len

	n := l.head
	for j := 0; j < i; j++ {
		n = n.Next()
	}
	return n
}
