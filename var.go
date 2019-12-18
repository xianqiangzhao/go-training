package main

import "fmt"

var (
	pi = 3.14
)

func definevar() {
	var i int = 33
	var a, b, c, d = 1, 2, 3, "aaa"
	fmt.Println(i, a, b, c, d, pi)
	fmt.Printf("%d %d %d %s %f \n", a, b, c, d, pi)

}

func constvar() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func getGrade(score int) string {
	result := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score %d", score))
	case score < 60:
		result = "F"
	case score < 80:
		result = "C"
	case score < 90:
		result = "B"
	case score <= 100:
		result = "A"
	}
	return result
}
func main() {
	fmt.Printf("hello, world\n")
	definevar()
	constvar()
	fmt.Println(getGrade(10), getGrade(30), getGrade(68), getGrade(80))
}
