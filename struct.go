package main

import "fmt"

type tree struct {
    X int
    Y int
    left,right *tree
}
func (node *tree) print(){
    fmt.Println(node.X,node.Y)
    node.X = 250
    fmt.Println(node.X,node.Y)
}

func main() {
    v := tree{X:100, Y:200 }
    v.print()
}
