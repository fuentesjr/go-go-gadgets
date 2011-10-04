package main

import "testing"

func AssertLength(t *testing.T, list *LinkedListNode, expected_val int) {
  if list.Length() != expected_val {
    t.Errorf("Expected length to equal %d, but got %d", expected_val, list.Length())
  }
}

func TestLength(t *testing.T) {
  x := NewListNode(1)
  AssertLength(t, x, 1)
  x.Add(6)
  AssertLength(t, x, 2)
}
