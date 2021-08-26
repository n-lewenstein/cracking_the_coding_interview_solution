package main

/*
Algorith:
1. remove spaces from string
2. sort the string
3.
*/
import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "redder"
	fmt.Println("The string: " + str)
	str = removeSpace(str)
	fmt.Println("without spaces: " + str)
	str = sortString(str)
	fmt.Println("sorted: " + str)
	b := checkPalindrom(str)
	fmt.Println("Is a premutaion of palindrome? ")
	fmt.Println(b)

}

func removeSpace(s string) string {
	var b bytes.Buffer
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			b.WriteString(string(s[i]))
		}
	}
	return b.String()
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func checkPalindrom(w string) bool {

	nonPair := 0

	i := 0
	for i < len(w) && len(w) > 1 {
		if i == len(w)-1 {
			nonPair++
			break
		}
		if w[i] == w[i+1] {
			i += 2
		} else {
			i += 1
			nonPair++
		}
	}

	if nonPair < 2 {
		return true
	}
	return false
}
