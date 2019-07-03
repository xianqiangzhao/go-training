package main
import (
     "fmt"
)

func test() *int{
    x := 100
    return &x
}

func ftest(fn func() int) int {
return fn()
}

func main() {
    a := []int{0,0,0}
    b := make([]int, 3)
    c := new([]int)
    b[0] = 10
    fmt.Println(a, b ,c)
    s := "abcd"
    bs := []byte(s)
    bs[1] = 'B'
    fmt.Println(s, string(bs))
    for i := range s{
        fmt.Println(string(s[i]))
    }

    call_test := test()
    fmt.Printf("call_test= %d \n" , (*call_test) +1 )
    m := map[string]int{"a" : 1, "b": 2, "c":3}
    for k , v := range m {
        fmt.Printf("key = %s value = %d \n", k, v)
    }
    fmt.Println(len(m))
    s1 := ftest(func() int {return 100})
    fmt.Printf("call ftest= %d", s1)

}