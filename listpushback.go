package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		listItem := l.Head
		for listItem.Next != nil {
			listItem = listItem.Next
		}
		listItem.Next = n
		l.Tail = n
	}
}
