new to new 
111
>>>>>>> 16ac78e21dff426dbeedc600dd38817ba4ada2bc
old
<<<<<<< HEAD

=======
news
<<<<<<< HEAD
>>>>>>> cc641962a7f45d6ba3b99a16549d08c4f253d36b
aaa
=======
ddd
>>>>>>> b9189eadcef5fb26126f908a8f912ee35dce9fe9
package main
bbb
import "fmt"

func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	s = append(s, 2, 3, 4)
	s = append(s, 2, 3, 4)
	printSlice(s)
	var i = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	is := i[:]
	printSlice(is)
	is1 := is[0:4]
	printSlice(is1)
	is1 = append(is1, 100)
	printSlice(is1)
	printSlice(is)
	is1 = append(is1, 200)
	is1 = append(is1, 300)
	is1 = append(is1, 400)
	is1 = append(is1, 500)
	is1 = append(is1, 600)
	printSlice(is1)
	printSlice(is)
	var ss []int
	printSlice(ss)
	if ss == nil {
		fmt.Println("ss is nil")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
