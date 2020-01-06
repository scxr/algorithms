package main

import "fmt"

func combsort(collection []int) []int {
	shrinkfac := 1.3
	gap := len(collection)
	swapped := true
	i := 0

	for ; i < gap || swapped == true; i++ {
		swapped = false
		i = 0
		gap = int(float64(gap) / shrinkfac)
		for ; gap+i < len(collection); i++ {
			if collection[i] > collection[i+gap] {
				collection[i], collection[i+gap] = collection[i+gap], collection[i]
			}
		}
	}
	return collection
}

func main() {
	unsorted := []int{23, 3, 23, 2, 32, 3232, 3, 667567, 4, 436, 87}
	fmt.Println(combsort(unsorted))
}
