package main

import (
	"fmt"
	"strings"
)

var morsevals = map[string]string{ // dictionary of all morsevalues
	"A":  ".-",
	"B":  "-...",
	"C":  "-.-.",
	"D":  "-..",
	"E":  ".",
	"F":  "..-.",
	"G":  "--.",
	"H":  "....",
	"I":  "..",
	"J":  ".---",
	"K":  "-.-",
	"L":  ".-..",
	"M":  "--",
	"N":  "-.",
	"O":  "---",
	"P":  ".--.",
	"Q":  "--.-",
	"R":  ".-.",
	"S":  "...",
	"T":  "-",
	"U":  "..-",
	"V":  "...-",
	"W":  ".--",
	"X":  "-..-",
	"Y":  "-.--",
	"Z":  "--..",
	"1":  ".----",
	"2":  "..---",
	"3":  "...--",
	"4":  "....-",
	"5":  ".....",
	"6":  "-....",
	"7":  "--...",
	"8":  "---..",
	"9":  "----.",
	"0":  "-----",
	", ": "--..--",
	".":  ".-.-.-",
	"?":  "..--..",
	"/":  "-..-.",
	"-":  "-....-",
	"(":  "-.--.",
	")":  "-.--.-",
}

func morseit(mystr string) string {
	decval := ""
	for _, char := range strings.ToUpper(mystr) { // convert it to upper and loop through
		decval += morsevals[string(char)] + " " // map the char to morse
	}
	return decval
}

//test
func main() {
	fmt.Println(morseit("oog"))
}
