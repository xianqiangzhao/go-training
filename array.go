package main
import "fmt"

func main() {
    var arr1 [5]int
    var a bool = true 
    arr2 := [3] int{1,2,3}
    arr3 := [...] int{3,4,5}
    var grid [4][5] int
    fmt.Println(arr1, arr2, arr3)
    fmt.Println(grid)
    for i := 0; i < len(arr3); i++ {
        fmt.Println(arr3[i])
    }
    if a {
        fmt.Println(arr3)
    }
    for _, v := range arr3 {
        fmt.Println(v)
    }
}