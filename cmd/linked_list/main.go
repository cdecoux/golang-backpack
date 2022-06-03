package main

import log "github.com/sirupsen/logrus"

type Node struct {
	Next *Node
	Value any
}

func (n *Node) ToList() []any {
	if n.Next == nil {
		return []any{n.Value}
	}
	return append([]any{n.Value}, n.Next.ToList()...)
}

func (n *Node) LastValueK(k int) any  {
	list := n.ToList()

	if k >= len(list) || k < 0 {
		return -1
	}
	return list[len(list) - (1 + k)]
}

func LinkedList[T any](list []T) *Node {

	var prev *Node
	for i := len(list) - 1; i >= 0; i-- {
		node := &Node{
			Next: prev,
			Value: list[i],
		}
		prev = node
	}
	return prev
}

func main() {

	l1 := LinkedList([]int{1, 2, 3})
	l2 := LinkedList([]string{"a", "b", "c", "d", "e"})
	l3 := LinkedList([]any{1, "a", true})

	log.Info(l1.ToList())
	log.Info(l2.ToList())
	log.Info(l3.ToList())

	log.Info(l1.LastValueK(0))
}