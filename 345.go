func reverseVowels(s string) string {
	bytes := []byte(s)
    left, right := 0, len(s) - 1

	for left < right {
		if !isVowel(bytes[left]) {
			left++
		}

		if !isVowel(bytes[right]) {
			right--
		}

		if isVowel(bytes[left]) && isVowel(bytes[right]) {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}

	return string(bytes)
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

/*
Input: s = "IceCreAm"

Output: "AceCreIm"

// move two pointers from the start and end of the string
// if the characters at the pointers are vowels, swap them
// if the characters at the pointers are not vowels, move the pointers inward
// continue until the pointers meet in the middle
// return the modified string
