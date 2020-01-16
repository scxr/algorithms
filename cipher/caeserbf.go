package main

import (
	"fmt"
	"strings"
)

func bruteforce(s string) string{
	var result string
	var retval string
	for i := 0; i <= 26;i++{
		result =""
		for _, char := range strings.ToLower(s){
			if int(char) == 32 {
				result += " " // 32 is space val
			} else if int(char) + i > 122 { // z is 122 ascii
				tmp := (int(char)+i) -122 
				if tmp == 26{ // 97 + 26 == 123 123>122 
					result += "a"
				} else {
					result += string(97 + tmp) // 97 is first ascii
				}

			} else {
				result += string(int(char) + i) 
				
			}
					
		}
		retval += result + "\n"
	}
	return retval
}

func main(){
	
	fmt.Println(bruteforce("khoor zruog")) // hello world (shift 3)
}
