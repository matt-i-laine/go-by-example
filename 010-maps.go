package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs1 := m["k1"]
	_, prs2 := m["k2"]
	fmt.Println("A value is present for key \"k1\"? ", prs1)
	fmt.Println("A value is present for key \"k2\"? ", prs2)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
