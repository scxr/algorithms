/*
A program to find exact fibonacci numbers in O(1) time, after n=76 however an integer overflow occurs
Based on https://i.stack.imgur.com/DjBGd.png
Im not sure on the math.Pow implementation in go though so it may be logn
*/
package main

import (
	"fmt"
	"math"
)

func fib(n uint64) uint64 {
	var phi float64 = (1 + math.Sqrt(float64(5))) / 2 // 
	return uint64(math.Round(math.Pow(phi, float64(n)) / math.Sqrt(float64(5)))) 
}

func main() {
	fmt.Println(fib(77))
}
