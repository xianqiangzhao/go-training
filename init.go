package main

import (
    "fmt"
    "time"
)
var now = time.Now()
var v = make(chan bool)
func main() {
    fmt.Println("main1:" , int(time.Now().Sub(now).Seconds()))
     <- v //wait  for init go
    fmt.Println("main2:" , int(time.Now().Sub(now).Seconds()))

}

func init() {
    fmt.Println("init:" , int(time.Now().Sub(now).Seconds()))
    go func() {
        time.Sleep(time.Second * 5)
        v <- true
    }()
}