package array

//5. 最长回文子串-https://leetcode.cn/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {

	l, r := 0, 0
	for i := 0; i < len(s); i++ {

		//1. 长度为奇数
		l1, r1 := expandAroundCenter(s, i, i)
		//2. 长度为偶数
		l2, r2 := expandAroundCenter(s, i, i+1)

		if r-l < r1-l1 {
			l = l1
			r = r1
		}

		if r-l < r2-l2 {
			l = l2
			r = r2
		}
	}
	return s[l : r+1]
}

func expandAroundCenter(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}
	return l + 1, r - 1
}
