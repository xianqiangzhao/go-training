package main

import (
    "fmt"
    "time"
    "runtime"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(1000 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    //go say("hello")
    //say("world")
    var a [10]int
    for i := 0; i < 10; i++ {
        go func(i int){
            for {
                //fmt.Printf("Hello from " +  "goroutines %d \n ", i)
                a[i]++
                runtime.Gosched()
            }
        }(i)
    }
    time.Sleep(time.Millisecond)
    fmt.Println(a)
    
}