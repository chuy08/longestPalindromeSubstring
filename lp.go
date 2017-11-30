package main

import "fmt"

func longestString(s string) {
	n := len(s)
	//fmt.Println(n)
	palindromeBeginsAt := 0
	maxLen := 1
	palindrome := make([][]bool, n)

	for i := range palindrome {
		palindrome[i] = make([]bool, n)
	}

	// Single letter palindrome
	for i := 0; i < n; i++ {
		palindrome[i][i] = true
	}

	// Double letter palindrome
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			palindrome[i][i+1] = true
			//palindromeBeginsAt = i
			//maxLen = 2
		}
	}

	for currLength := 3; currLength <= n; currLength++ {
		for i := 0; i < n-currLength+1; i++ {
			j := i + currLength - 1
			if s[i] == s[j] && palindrome[i+1][j-1] {
				palindrome[i][j] = true
				palindromeBeginsAt = i
				maxLen = currLength
			}
		}
	}
	//fmt.Println(palindrome)
	fmt.Println(palindromeBeginsAt, maxLen+palindromeBeginsAt)
}

func main() {
	s := "banana"
	longestString(s)
}
