package array

//344. 反转字符串 -https://leetcode.cn/problems/reverse-string/
func reverseString(s []byte) {

	//0.head tail
	head := 0
	tail := len(s) - 1

	//1. reverse
	for head < tail {
		s[head], s[tail] = s[tail], s[head]
		head++
		tail--
	}
	return
}
