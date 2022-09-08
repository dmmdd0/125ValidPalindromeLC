package main

import (
	"fmt"
)

func main() {

	palindrom := ""
	palindrom = "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(palindrom))

}

func isPalindrome(s string) bool {

	// remove all no letter & lower case
	res := make([]string, 0, len(s))
	for _, v := range s {
		switch {
		case v >= 90 && v <= 122:
			res = append(res, string(v))
		case v >= 65 && v <= 90:
			res = append(res, string(v+32))
		}
	}

	return false
}
