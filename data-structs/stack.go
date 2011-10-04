package main

type Stack struct {
  buffer []int
  index_tip int
}

func (s *Stack) Push(data int) {
  s.buffer = append(s.buffer, data)
}

func (s *Stack) Pop() {
}

func NewStack() (*Stack){
  node := new(Stack)
  return node
}
