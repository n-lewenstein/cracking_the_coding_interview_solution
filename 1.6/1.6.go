package main

/*
Algorith:
1. remove spaces from string
2. sort the string
3.
*/
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(stringCompression("aaabbcccceedddd"))
}

func stringCompression(s string) string {
	var b byte = s[0]
	counter := 0
	var newString strings.Builder
	newString.WriteString("")
	for i := 0; i < len(s); i++ {
		if s[i] == b && i != len(s)-1 {
			counter++
		} else if i == len(s)-1 {
			newString.WriteByte(b)
			newString.WriteString(fmt.Sprint(counter))
			counter = 1
			b = s[i]
		} else {
			if s[i] == b {
				newString.WriteByte(b)
				newString.WriteString(fmt.Sprint(counter + 1))
			} else {
				newString.WriteByte(b)
				newString.WriteString(fmt.Sprint(counter))
				newString.WriteByte(s[i])
				newString.WriteString(fmt.Sprint(1))
			}
		}
	}
	if newString.String() != "" {
		return newString.String()
	} else {
		return s
	}
}
