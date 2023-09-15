package main

import (
	"fmt"
	"time"
)

func doSomething(size int, c chan int) {
    for i := 0; i < size; i++ {
        time.Sleep(100 * time.Millisecond)
    }
    c <- size
}

func doAll(c chan int) {
    d:= make(chan int)
    go doSomething(10, d)
    go doSomething(20, d)
    go doSomething(30, d)
    c <- (<-d)
    c <- (<-d)
    c <- (<-d)
    close(c)
}

func main() {
    c := make(chan int)

    go doAll(c)
    for i := range c {
        fmt.Println(i)
    }
}

// 10
// 20
// 30
