package palindrome_number

import (
	"testing"
)

func TestFunc(t *testing.T) {
	if !isPalindrome(121) {
		t.Error("faile")
	}
	if isPalindrome(-121) {
		t.Error("faile")
	}

	if !isPalindrome(11) {
		t.Error("faile")
	}

	if isPalindrome(12) {
		t.Error("faile")
	}

	if !isPalindrome(+11) {
		t.Error("faile")
	}
}
