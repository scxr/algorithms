package main
/*
import (
	"fmt"
)
*/
func main() {
        unsorted := []int{-1,2,0,1,0,3,3,13,28,4}
        QuickSort(unsorted)
	
}


func QuickSort(unsorted []int) {
	if len(unsorted) <= 1 {
		return
	} // find the best pivot

	i := Partition(unsorted)
	QuickSort(unsorted[0:i])
	QuickSort(unsorted[i+1:])
}

//Hoare partition
func Partition(unsorted []int) int {
	if len(unsorted) <= 1 {
		return 0
	}

	pivot := unsorted[0]
	i, j := 0, len(unsorted)-1

	for {
		for i < len(unsorted) && unsorted[i] <= pivot {
			i++
		}

		for unsorted[j] > pivot {
			j--
		}

		if i >= j { // if it is greater than then swap the elements
			unsorted[0], unsorted[j] = unsorted[j], unsorted[0]
			return j
		}

		unsorted[i], unsorted[j] = unsorted[j], unsorted[i] // swap the elements
	}
}
