func mergeAlternately(word1 string, word2 string) string {
	var result []byte

	i, j := 0, 0
	n, m := len(word1), len(word2)

	for i < n && j < m {
		result = append(result, word1[i], word2[j])
		i++
		j++
	}

	if i < n {
		result = append(result,word1[i:]...)
	}

	if j < m {
		result = append(result, word2[j:]...)
	}

	return string(result)
}


/*
time: O(n+m)
space: O(n+m)
*/

/*
Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"

Vì:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r

word1:  a   b 
word2:    p   q   r   s
merged: a p b q   r   s