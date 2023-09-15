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

func main() {
    c, q := make(chan int), make(chan int)
    jobs := 5

    go func() {
        for i := 1; i <= jobs; i++ {
            doSomething(i*10, c)
        }
        q <- 0 // done
    }()

    for {
        select {
        case x := <-c: // if we have a result
            fmt.Println(x)
        case <-q: // if we are done
            fmt.Println("Finished")
            return
        default: // if we are waiting. The default case will run if no other channel is ready.
            fmt.Print("...")
            time.Sleep(time.Second)
        }
    }
}


// ......10
// .........20
// ............30
// ............40
// ..................50
// Finished