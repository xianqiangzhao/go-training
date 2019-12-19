package main

import "fmt"

const p = "death & taxes"
const (
	pi       = 3.14
	language = "Go"
)

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println(pi)
}
