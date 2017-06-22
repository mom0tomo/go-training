package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4}

	for v := range arr {
		fmt.Println(v)
	}
}
