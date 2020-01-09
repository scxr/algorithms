package main

import "fmt"

func caeserenc(mystr string, shiftnum int) string {
	out := ""
	for _, char := range mystr {
		intchar := int(char)
		if 65 <= intchar && intchar <= 90 { //uppercase characters
			out += string((int(char)+shiftnum-65)%26 + 65)
		} else if 97 <= intchar && intchar <= 122 { //lowercase characters
			out += string((int(char)+shiftnum-97)%26 + 97)
		} else {
			out += string(char) //for example numbers
		}

	}
	return out
}

func main() {
	fmt.Println(caeserenc("hello world", 1))
}
