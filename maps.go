package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

type person struct {
    name string, 
    height int
}

var m map[string]Vertex
var k map[string]person

func main() {
    m = make(map[string]Vertex, 5)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    m["22Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m, len(m))

    k = make(map[string]person)
    k["a"] = person{"zhao" }
    k["b"] = person{"laoer" }


}
