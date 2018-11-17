package palindrome_number

import (
	"strconv"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)

	length := len(str) - 1

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[length-i] {
			return false
		}
	}
	return true
}
