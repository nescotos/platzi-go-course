package main

import "fmt"

func main() {
	m1 := make(map[string]int)

	m1["a"] = 8

	m1[9] = "asas"

	fmt.Println(m1)
	fmt.Println(m1["a"])
}
