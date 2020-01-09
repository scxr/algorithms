package main

func gcd(a, b int) int {
	for a > 0 {
		a, b = b%a, a // set a to b mod a and b to a, we mod until we find the gcd
	}
	return b
}

func fmi(a, b int) int {
	if gcd(a, b) != 1 { // only works for prime numbers
		return 0
	}
	u1, u2, u3 := 1, 0, a
	v1, v2, v3 := 0, 1, b

	for v3 != 0 {
		x := int(u1 / v3)
		v1, v2, v3, u1, u2, u3 = (u1 - x*v1), (u2 - x*v2), (u3 - x*v3), v1, v2, v3
	}
	return u1 % b
}
