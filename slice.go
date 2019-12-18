package main

import "fmt"

func update_arr(a [6]int, up int) {
	a[2] = up
}

func update_slice(a []int, up int) {
	a[2] = up
}

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	update_arr(primes, 200)
	fmt.Println(primes)
	update_slice(s, 300)
	fmt.Println(len(s), cap(s), s)
	fmt.Println(len(primes), cap(primes), primes)
}

/* result
[2 3 5 7 11 13]
3 5 [3 5 300]
6 6 [2 3 5 300 11 13]


切片就像数组的引用
切片并不存储任何数据，它只是描述了底层数组中的一段。
更改切片的元素会修改其底层数组中对应的元素。
与它共享底层数组的切片都会观测到这些修改。
*/
