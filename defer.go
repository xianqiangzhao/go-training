package main

import "fmt"

func f(result int) int {
	defer func() {
		result += 1
	}()
	return result
}
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println(f(11))

}
