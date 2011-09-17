package main

import (
    "os"
    "flag"  // command line option parser
    //"strconv"
    "strings"
)


var omitNewline = flag.Bool("n", false, "don't print final newline")


const (
    Space = " "
    Newline = "\n"
)


func main() {
    flag.Parse()   // Scans the arg list and sets up flags
    s := ""
    numArgs := flag.NArg()

    os.Stdout.WriteString(strings.Join(flag.Args(), "-") + "\n")
    for i := 0; i < numArgs; i++ {
      //os.Stdout.WriteString(strconv.Itoa(numArgs) + "\n")
        if i > 0 {
            s += Space
        }
        s += flag.Arg(i)
    }
    if !*omitNewline {
        s += Newline
    }
    os.Stdout.WriteString(s)
}
