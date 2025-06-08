func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLength := gcd(len(str1), len(str2))
	return str1[:gcdLength]
}	

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

/*
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"

Vì:
str1:  A   B   C   A   B   C
str2:    A   B   C

*/