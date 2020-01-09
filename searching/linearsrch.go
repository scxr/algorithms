package main

import "fmt"

func linsrch(srchval int, srcharr []int) string {
	for index, val := range srcharr { //iterate through the elements
		if val == srchval { // if its equal then search successful
			return fmt.Sprintf("%s%d", "found at index ", index) // return it
		}
	}
	return "not found" // the value hasnt been found if we reach here
}

//tests
func main() {
	srcharr := []int{1, 2, 3, 4, 7, 8, 10, 454, 455, 666, 3333, 4444}
	fmt.Println(linsrch(666, srcharr))
}
