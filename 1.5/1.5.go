package main

/*
Algorith:
1. If the absolute difference between array length >= 2 return false
2. sort the string
3.
*/
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(oneAway("pale", "ple"))
	fmt.Println(oneAway("pales", "pale"))
	fmt.Println(oneAway("pale", "bale"))
	fmt.Println(oneAway("pale", "bake"))
}

func oneAway(s1, s2 string) bool {
	length := 0
	if len(s1) > len(s2) {
		length = len(s2)
	} else {
		length = len(s1)
	}

	if math.Abs((float64)(len(s1)-len(s2))) >= 2 {
		return false
	}

	i, j := 0, 0
	counter := 0

	for i < length && j < length {
		if counter >= 2 {
			return false
		}
		if s1[i] == s2[j] {
			i += 1
			j += 1
			continue
		}
		if j+1 < len(s2) && s1[i] == s2[j+1] {
			j += 1
			counter += 1
			continue
		}
		if i+1 < len(s1) && s1[i+1] == s2[j] {
			i += 1
			counter += 1
			continue
		}
		i += 1
		j += 1
		counter += 1
	}
	return true
}
