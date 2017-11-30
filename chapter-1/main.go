package main

import (
    "os"
    "fmt"
)

func getFormat() (string, bool) {
    return "%s %s\n", true
}

func main() {
    if len(os.Args) != 2 {
        println("Needs one arg")
        os.Exit(1);
    }
    msg, name := "Hello", os.Args[1]
    format, ok := getFormat()
    if !ok {
        println("Not OK")
        os.Exit(1);
    }
    fmt.Printf(format, msg, name)
}
