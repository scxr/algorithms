package main

import "fmt"

func bubblesort(collection []int) []int {

	for i := 0; i < len(collection)-1; i++ {
		fmt.Println("made it here")
		swapped := true
		for x := 0; x < len(collection)-1-i; x++ {
			if collection[x] > collection[x+1] {
				swapped = true
				collection[x], collection[x+1] = collection[x+1], collection[x]
			}

			if swapped == false {
				return collection
			}

		}
	}
	return collection
}

func main() {
	unsorted := []int{23, 323, 2, 1, 6, 7, 4, 6, 7}
	fmt.Println(bubblesort(unsorted))
}
