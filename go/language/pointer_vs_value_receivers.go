package main

import "fmt"

var counter int

type t int
type p struct {
	*t
}

type doer interface {
	do()
}

func dodo(d doer) {
	d.do()
}

func (a t) do() {
	a = t(counter)
	counter++
	fmt.Println("in do a equals ", a)
}

func main() {
	var r t
	var y p
	y.t = &r
	dodo(y)

	var z = &y
	dodo(z)
}
