package main

import "fmt"

func decodeString(s string) string {
	countStack := make([]int, 0)
	strStack := make([]string, 0)
	currentNum := 0
	currentStr := ""

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= '0' && ch <= '9' {
			currentNum = currentNum*10 + int(ch-'0')
		} else if ch == '[' {
			countStack = append(countStack, currentNum)
			strStack = append(strStack, currentStr)

			currentNum = 0
			currentStr = ""
		} else if ch == ']' {
			repeatTimes := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]
			
			prevStr := strStack[len(strStack) - 1]
			strStack = strStack[:len(strStack)-1]

			repeated := ""
			for j := 0; j < repeatTimes; j++ {
                repeated += currentStr
            }

			currentStr = prevStr + repeated

		} else {
			currentStr += string(ch)
		}

	}

	return currentStr

	
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
}


/*
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 105.

 

Example 1:

Input: s = "3[a]2[bc]"
Output: "aaabcbc"
Example 2:

Input: s = "3[a2[c]]"
Output: "accaccacc"
Example 3:

Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
*/