package main

import (
	"fmt"
	"math"
)

func mp(srcharr []int) (int, int) {
	r := len(srcharr) - 1
	l := 0
	index := int((math.Floor(float64((l + r) / 2)))) // we get the floor of the middle value
	return index, srcharr[index]                     // return mp index and mp value
}

func binsrch(srchval int, arr []int) string {
	found := false
	for found == false {
		index, currval := mp(arr)
		if srchval == currval {
			return fmt.Sprintf("%s%d%s", "the value ", srchval, " has been found ") // return where the value was found as our mp was equal to our srch val
		} else if len(arr) <= 1 { //if our array is lenght of 1 or less and wasnt found by the if above then it isnt in the array
			return "couldnt be found"
		} else if srchval < currval {
			arr = arr[:index] // erase right side of the array
		} else if srchval > currval {
			arr = arr[index:] // erase left side of array
		} else {
			return "value couldnt be found" // something went wrong
		}
	}
	return "error somewhere"
}

// just testing it
func main() {
	srcharr := []int{1, 2, 3, 4, 7, 8, 10, 454, 455, 666, 3333, 4444}
	fmt.Println(binsrch(656, srcharr))
}
