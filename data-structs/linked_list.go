package main

type LinkedListNode struct {
  value int
  next *LinkedListNode
}

func NewListNode(val int) (*LinkedListNode){
  node := new(LinkedListNode)
  node.value = val
  node.next = nil

  return node
}

func (list *LinkedListNode) Add(val int) {
  newnode := NewListNode(val)
  if list.next != nil {
    newnode.next = list.next
  }
  list.next = newnode
}

func (list *LinkedListNode) Length() (int) {
  count := 1
  for node := list.next; node != nil; node = node.next {
    count++
  }

  return count
}
