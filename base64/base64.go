package main
import (
  "fmt"
  "io"
  "encoding/base64"
  "os"
)


func main() {
  fmt.Printf("Hello World\n")
  myfile := os.OpenFile("blah.txt")
  myimgfile =: os.OpenFile("voltron.jpg")
  writter := NewEncoder(encoding.base64.StdEncoding, myfile)
  myimgfile.Read()
  //myimgfile.Stat().Size
}
