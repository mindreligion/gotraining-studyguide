package main

import "fmt"

func main() {
	s := [5]string{"a", "b", "c", "d", "e"}

	for i, v := range s {
		s[1] = "zalupa"

		if i == 0 {
			fmt.Printf("val sem 0 index addr is %p\n", &s[i])
		}
		fmt.Println(v)
	}

	for i := range s {
		s[1] = "zalupa"
		if i == 0 {
			fmt.Printf("point sem 0 index addr is %p\n", &s[i])
		}
		fmt.Println(s[i])
	}

	fmt.Printf("\n\n\n\n\n\n\n")

}
