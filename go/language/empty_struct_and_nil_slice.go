package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a struct{}
	var b struct{}

	fmt.Printf("a value is %v, a address is %p\n", a, &a)
	fmt.Printf("b value is %v, b address is %p\n", b, &b)

	var s []bool = []bool{}

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	dataPtr := (*[0]bool)(unsafe.Pointer(hdr.Data))

	fmt.Printf("%+v\n", hdr)
	fmt.Printf("value of a ptr %p\n", dataPtr)
	fmt.Printf("value of a struct{} %+v\n", dataPtr)

	var arr [0]int
	fmt.Printf("addr of empty array %p\n", &arr)

	_ = append(s, false)

	str := "рот ебал, how are you?"
	fmt.Printf("str[0] type is %T, value is %v\n", str[0], str[:1])
	for i, r := range str {
		fmt.Println(string(r))
		fmt.Printf("i on current iteration is %v\n", i)
	}

}
