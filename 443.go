package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	write := 0
	read := 0

	for read < len(chars) {
		currentChar := chars[read]
		count := 0
		for read < len(chars) && chars[read] == currentChar {
			read++
			count++
		}

		chars[write] = currentChar
		write++

		if count > 1 {
			for _, digit := range []byte(strconv.Itoa(count)) {
				chars[write] = digit
				write++
			}
		}

		fmt.Println(chars)
	}

	return write
}

func main() {
	compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'})
}

/*

Bạn có một mảng ký tự chars (ví dụ: ['a','a','b','b','c','c','c']).
Bạn cần nén mảng này theo nhóm ký tự lặp liên tiếp, bằng cách:
Với mỗi nhóm ký tự giống nhau liên tiếp (ví dụ nhóm 'a','a' hay 'c','c','c'):
Nếu nhóm chỉ có 1 ký tự, giữ nguyên ký tự đó (không thêm số).
Nếu nhóm có độ dài > 1, bạn ghi ký tự đó, sau đó ghi độ dài nhóm dưới dạng chuỗi ký tự (ví dụ '3' cho nhóm 3 ký tự).
Mảng chars sẽ được chỉnh sửa trực tiếp tại chỗ, không trả về một chuỗi mới mà thay đổi mảng input luôn.
Cuối cùng, bạn trả về độ dài mới của mảng sau khi nén (không phải mảng, mà số nguyên).

Example 1:

Input: chars = ["a","a","b","b","c","c","c"]
Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".
*/
