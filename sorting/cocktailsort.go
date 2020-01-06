package main

import "fmt"

func cocktailsort(collection []int) []int {
	unsortlen := len(collection) - 1
	swapped := false
	for i := unsortlen; unsortlen > 0; unsortlen-- {

		for j := i; j > 0; j-- {
			if collection[j] < collection[j-1] {
				collection[j], collection[j-1] = collection[j-1], collection[j]
				swapped = true
			}
		}
		for j := 0; j < i; j++ {
			if collection[j] > collection[j+1] {
				collection[j], collection[j+1] = collection[j+1], collection[j]
				swapped = true
			}
		}
		if swapped == false {
			return collection
		}
	}
	return collection
}

func main() {
	unsorted := []int{32, 534, 2, 222, 5, 3, 2}
	fmt.Println(cocktailsort(unsorted))
}
