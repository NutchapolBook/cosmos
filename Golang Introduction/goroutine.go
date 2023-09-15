package main

import (
	"fmt"
	"time"
)

func doSomething(size int) {
    for i := 0; i < size; i++ {
        fmt.Println(i)
        time.Sleep(time.Second)
    }
}

func main() {
    go doSomething(10) // go statement before a function creates a goroutine
    go doSomething(10)
    time.Sleep(10*time.Second)
}

// 0
// 0
// 1
// 1
// 2
// 2
// 3
// 3
// 4
// 4
// 5
// 5
// 6
// 6
// 7
// 7
// 8
// 8
// 9
// 9