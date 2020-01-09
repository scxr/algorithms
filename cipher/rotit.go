package main

import "fmt"

func rot(mystr string, rotval int) string {
	out := ""
	for _, char := range mystr { // ignore the index we dont need it 
		if int(char) >= 65 && int(char) <= 90 { // if it us an upper case alpha character 
			out += string(65 + (int(char)-65+rotval)%26) // rotate it by rotate value
		} else if int(char) >= 97 && int(char) <= 122 { // if it is a lowecase character
			out += string(97 + (int(char)-97+rotval)%26) // rotate it by rotate value 
		} else {
			out += string(char) // if it isnt a alpha characters
		}

	}
	return out

}
//testing
func main() {
	fmt.Println(rot("hello", 13))
}
