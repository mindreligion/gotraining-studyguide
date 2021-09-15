package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := make([]int, 3)
	c := make([]int, 5)

	l := copy(b, a)
	fmt.Println(b, l)
	l = copy(c, a)
	fmt.Println(c, l)
	cc := c[2:]
	l = copy(cc, c)
	fmt.Println(cc, l)
	fmt.Println(c)
	l = copy(cc, c)
	fmt.Println(cc, l)
	fmt.Println(c)
	l = copy(cc, c)
	fmt.Println(cc, l)
	fmt.Println(c)
	l = copy(cc, c)
	fmt.Println(cc, l)
	fmt.Println(c)

	st := "sobaka"
	bb := make([]byte, len(st))
	l = copy(bb, st)
	fmt.Println(string(bb))
}
