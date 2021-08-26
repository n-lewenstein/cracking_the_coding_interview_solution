package main

/*
Algorith:
1. function ispermutation(string1, string2)
2. if length of both strings are not equal return false(can't be a premutation)
3. sort both strings
4. compare both string and return
*/

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()

	//s1 := []rune{'a', 'b', 'c', 'd'}
	//s2 := []rune{'d', 'c', 'b', 'a'}
	s1 := []string{"a", "b", "c", "d"}
	s2 := []string{"f", "c", "b", "d"}
	fmt.Println(is_premutation(s1, s2))

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s", elapsed)
}

func is_premutation(s1, s2 []string) bool {

	if len(s1) != len(s2) {
		return false
	}

	sort.Strings(s1)
	sort.Strings(s2)

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
