package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Run("'kayak' is a palindrome", func(t *testing.T) {
		s := "kayak"
		actual := IsPalindrome(s)
		expected := true

		if actual != expected {
			t.Errorf("actual %t expected %t, given %q", actual, expected, s)
		}

	})

	t.Run("error is not a palindrome", func(t *testing.T) {
		s := "error"
		actual := IsPalindrome(s)
		expected := false

		if actual != expected {
			t.Errorf("actual %t expected %t, givn %q", actual, expected, s)
		}
	})

}
