package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	count1 := make(map[rune]int)
	count2 := make(map[rune]int)

	for _, v := range word1 {
		count1[v]++
	}
	for _, v := range word2 {
		count2[v]++
	}

	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)
	for ch := range count1 {
		set1[ch] = true
	}
	for ch := range count2 {
		set2[ch] = true
	}
	if len(set1) != len(set2) {
		return false
	}
	for ch := range set1 {
		if _, ok := set2[ch]; !ok {
			return false
		}
	}

	freq1 := make([]int, 0, len(count1))
	freq2 := make([]int, 0, len(count2))
	for _, v := range count1 {
		freq1 = append(freq1, v)
	}
	for _, v := range count2 {
		freq2 = append(freq2, v)
	}

	sort.Ints(freq1)
	sort.Ints(freq2)

	if len(freq1) != len(freq2) {
		return false
	}
	for i := 0; i < len(freq1); i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(closeStrings("abc", "bca"))
}

/*

Two strings are considered close if you can attain one from the other using the following operations:

Operation 1: Swap any two existing characters.
For example, abcde -> aecdb
Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
You can use the operations on either string as many times as necessary.

Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.

*/
