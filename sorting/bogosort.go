package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bogosort(collection []int) []int {
	for issorted(collection) == false {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(collection), func(i int, j int) { collection[i], collection[j] = collection[j], collection[i] })
	}
	return collection

}

func issorted(collection []int) bool {
	if len(collection) < 2 {
		return true
	}
	for i := 0; i < len(collection)-1; i++ {
		if collection[i] > collection[i+1] {
			return false
		}

	}
	return true
}

func main() {
	var unsorted = []int{3, 5, 213, 23, 2, 1, 66}
	fmt.Println(bogosort(unsorted))
}
