package main

import "fmt"

func main() {
	//a := make([]int, 5)
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	var p *[]int = new([]int)
	var v []int = make([]int, 10, 100)
	fmt.Println(p, v)
	if *p == nil {
		fmt.Println("p is nil")
	}
	if v == nil {
		fmt.Println("v is nil")
	}

}

func printSlice(s string, x []int) {
	x[0] = 100
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
