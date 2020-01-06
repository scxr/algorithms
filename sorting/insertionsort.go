package main

import (
	"fmt"
)

var insertindex int

func insertionsort(collection []int) []int {
	for loopindex := 1; loopindex < len(collection); loopindex++ {
		insertindex = loopindex
		for ; insertindex > 0 && collection[insertindex-1] > collection[insertindex]; insertindex-- {
			collection[insertindex] = collection[insertindex-1]
			collection[insertindex-1] = collection[insertindex]

		}
	}
	return collection
}

func main() {
	unsorted := []int{3, 4, 1, 5, 2, 232,2,432,1} // values to be sorted
	fmt.Println(insertionsort(unsorted))

}
