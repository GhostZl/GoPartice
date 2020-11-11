package main

import "fmt"

func changeSlice(ss []int) {
	ss[1] = 5
}
func testCopy() {
	fmt.Println("TESTING COPY METHED")
	a := []int{1, 2, 3, 4}
	var b = []int{3, 4}
	copy(b, a)
	fmt.Println(cap(a), len(b))
	fmt.Println(cap(b), len(b))
	fmt.Println(a, b)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(cap(a))
	a = append(a, 7)
	fmt.Println(cap(a))
	b := a[1:3]
	b[0] = 5
	c := append(b, 9)
	fmt.Println(a, b, c)
	testCopy()
}
