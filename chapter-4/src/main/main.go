package main

import (
    "shopping"
    "fmt"
)

type Logger interface {
    Log (msg string)
}

type ConsoleLogger struct {}

func (l ConsoleLogger) Log(mesg string) {
    fmt.Println(mesg)
}

func doTheLogThing(l Logger) {
    l.Log("hi")
    println(l)
}


func main() {
    fmt.Println(shopping.PriceCheck(4343))
    doTheLogThing(ConsoleLogger{})
}
