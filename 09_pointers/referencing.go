package main

import "fmt"

func zero(z *int) {
	fmt.Println(z)
	*z = 0
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x) // x is 0

	fmt.Println("================ \n")

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220

	var b = &a
	fmt.Println(b, &b) // 0x20818a220
	fmt.Println(*b)    // 43

}
