package main
import "fmt"

func main() {
  x := NewListNode(1)
  x.Add(2)
  x.Add(3)
  x.Add(4)
  fmt.Printf("length = %d\n", x.Length())
  y := NewStack()
  y.Push(7)
}
