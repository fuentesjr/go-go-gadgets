package main
//import "time"
import "fmt"
import "csv"
import "os"
import "regexp"

func exit_on_errors(err os.Error) {
  if err != nil && err != os.EOF {
    //log.Fatal(err)
    fmt.Printf("Received error: %s ... terminating with code 1.\n", err)
    os.Exit(1)
  }
}

func get_layout(text string) (string, os.Error) {
  // Map of format layouts to regular expressions
  regexps := make(map[string]*regexp.Regexp)
  regexps["01/02/2006 03:04"], _ = regexp.Compile("^[0-9][0-9]/[0-9][0-9]/([0-9][0-9]|[0-9][0-9][0-9][0-9])")
  regexps["01-02-2006 03:04"], _ = regexp.Compile("^[0-9][0-9]-[0-9][0-9]-[0-9][0-9][0-9][0-9]")
  regexps["01-02-06 03:04"], _ = regexp.Compile("^[0-9][0-9]-[0-9][0-9]-[0-9][0-9]")


  for layout, re := range regexps {
    //matches := re.Find(bytes.NewBufferString(text).Bytes()) 
    matches := re.FindString(text)
    if matches != "" {
      fmt.Println("Found match and layout ==> " + layout)
      return layout, nil
    }
  }

  return "", fmt.Errorf("Unrecognized pattern in: %g", text)
}

func main() {
  filename := "./raw_subscribers.csv"
  f, err := os.Open(filename)
  if err != nil { fmt.Println("Failed to open " + filename); os.Exit(1) }
  defer f.Close()

  filename = "test.csv"
  outf, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
  if err != nil { fmt.Println(err); os.Exit(1) }
  defer outf.Close()

  c := csv.NewReader(f)
  c.TrailingComma = true
  outc := csv.NewWriter(outf)
  defer outc.Flush()

  // Grab CSV headers
  _, err = c.Read()
  exit_on_errors(err)

  for {
    row, err := c.Read()
    if err != nil {
      if err != os.EOF {
        fmt.Println(err)
        os.Exit(1)
      } else {
        break;
      }
    }

    //layout, err := get_layout(row[0])
    //fmt.Println("Using layout ==> " + layout)
    //t, _ := time.Parse(layout, row[0])
    //fmt.Println(t)

    /*
    for i, field := range row {
      fmt.Printf(headers[i] + " => " + field + "\n")
    }
    */
    if err := outc.Write(row); err != nil {
      fmt.Println(err)
    }
  }

  fmt.Println("Done. Good bye.\n")
}
