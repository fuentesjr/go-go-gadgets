package main
import "time"
import "fmt"
import "csv"
import "os"

func exit_if_done(err os.Error) {
  if err != nil {
    if err == os.EOF {
      fmt.Printf("EOF.\n")
      os.Exit(0)
    } else {
      fmt.Printf("Received error... terminating.\n")
      os.Exit(1)
    }
  }
}

func main() {
  filename := "./data.csv"
  f, err := os.Open(filename)
  if err != nil { fmt.Println("Failed to open " + filename); os.Exit(1) }
  defer f.Close()

  c := csv.NewReader(f)
  layout := "01/02/2006"

  // Grab CSV headers
  headers, err := c.Read()
  exit_if_done(err)

  //for row, err := c.Read() && err != os.EOF {
  for row, err := c.Read();; row, err = c.Read() {
    exit_if_done(err)
    t, _ := time.Parse(layout, row[0])
    fmt.Println(t)

    for i, field := range row {
      fmt.Printf(headers[i] + " => " + field + "\n")
    }
  }
}
