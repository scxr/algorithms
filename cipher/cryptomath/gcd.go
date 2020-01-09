package main

import "fmt"

func gcd(a, b int) int {
	for a > 0 {
		a, b = b%a, a // set a to b mod a and b to a, we mod until we find the gcd
	}
	return b
}

//test
func main() {
	fmt.Println(gcd(4, 2))
}
