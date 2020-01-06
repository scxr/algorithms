package main

import "fmt"

func quicksort(coll []int) string {

	listlen := len(coll)

	if listlen <= 1 {
		return "already sorted"
	} else {
		finval := len(coll) - 1
		pivot := coll[finval]
		var greater = []int{}
		var lesser = []int{}
		for _, element := range coll {
			if element > pivot {
				greater = append(greater, element)
			} else {
				lesser = append(lesser, element)
			}

		}
		lessorted := quicksort(lesser)
		greatersorted := quicksort(greater)
		retval := fmt.Sprintf("%#v%d%#v", lessorted, pivot, greatersorted)
		return retval
	}
	

}

func main() {
	var arr = []int{23, 42, 2, 4, 2, 44, 5}
	fmt.Println(quicksort(arr))
}
