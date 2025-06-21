package main

import "fmt"

func maxVowels(s string, k int) int {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	maxVowels := 0
	sum := 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			sum++
		}
	}

	maxVowels = sum

	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			sum--
		}

		if isVowel(s[i]) {
			sum++
		}

		if sum > maxVowels {
			maxVowels = sum
		}
	}

	return maxVowels
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}

/*

Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.



Example 1:

Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.
Example 2:

Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.
Example 3:

Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.

*/
