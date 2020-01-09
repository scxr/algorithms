/*
Atbash is a substitution cipher where all the letters are revered. for example A->Z a->z b->y B->Y etc..
*/
package main

import "fmt"

func atbash(mystr string) string {

	out := ""
	for _, char := range mystr {
		intchar := int(char)
		if 65 <= intchar && intchar <= 90 { //upper case characters
			out += string(155 - intchar) // reverse
		} else if 97 <= intchar && intchar <= 122 { //lower characters
			out += string(219 - intchar) // reverse
		} else {
			out += string(char) //for example numbers
		}

	}
	return out
}
//test
func main() {
	fmt.Print(atbash("Test"))
}
