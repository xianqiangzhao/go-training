package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{100, 200}
    fmt.Println(Vertex{1, 2}, v.X)
}
