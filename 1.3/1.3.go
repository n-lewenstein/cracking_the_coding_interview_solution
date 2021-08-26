package main

/*
Algorith:
1. funnction recives string(str) and length(length)
2. returns a new string which replaces every space with %20
3. loop through str length stpets,
if a space appened %20 to new string else append str[i]
*/
import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	str := "hello space world   "
	fmt.Println(fill_spaces(str, 17))

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s", elapsed)
}

func fill_spaces(str string, length int) string {

	newString := ""
	astr := []rune(str)

	for i := 0; i < length; i++ {
		if astr[i] == ' ' {
			newString = newString + "%20"
		} else {
			newString = newString + string(str[i])
		}
	}
	return newString
}
