package main

import "fmt"

func lengthOfNonRepeatSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for index, ch := range []byte(s) {
		if _, ok := lastOccurred[ch]; ok && lastOccurred[ch] >= start {
			start = lastOccurred[ch] + 1
			length := index - start + 1
			maxLength = max(length, maxLength)
		} else {
			maxLength = index - start + 1
		}
		lastOccurred[ch] = index
	}
	return maxLength
}
func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
func main() {
	fmt.Println(lengthOfNonRepeatSubStr("nfpdmpi"))
	fmt.Println(lengthOfNonRepeatSubStr("bad"))
	s := "我爱慕课网"
	for i, c := range []rune(s) {
		fmt.Printf("(%d, %c)", i, c)
	}
}
