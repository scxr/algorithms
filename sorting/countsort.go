package main

import "fmt"

func countsort(collection []int) []int {
	if len(collection) <= 1 {
		return collection
	}
	collectionlen := len(collection)
	collmax, collmin := minmax(collection)
	countarrlen := collmax + 1 - collmin

	countarr := make([]int, countarrlen)
	ordered := make([]int, collectionlen)
	for index, _ := range countarr {
		countarr[index] = 0
	}

	for number := 0; number < len(collection); number++ {
		numval := collection[number]
		countarr[numval-collmin]++
	}

	for i := 1; i < countarrlen; i++ {
		countarr[i] = countarr[i] + countarr[i-1]
	}

	for index, _ := range ordered {
		ordered[index] = 0
	}

	for x := collectionlen - 1; x > -1; x-- {

		ordered[countarr[collection[x]-collmin]-1] = collection[x]
		countarr[collection[x]-collmin] -= 1
	}
	return ordered

}

func minmax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return max, min
}

func main() {
	unsorted := []int{323, 2, 3, 43, 6342, 6, 324, 6}
	fmt.Println(countsort(unsorted))
}
