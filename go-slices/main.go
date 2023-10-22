package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string

	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for _, x := range animals {
		fmt.Println(x)
	}

	fmt.Println("Element 0 is", animals[0])
	fmt.Println("First two elements are", animals[0:2])

	fmt.Println("The slice is", len(animals), "long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
