/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	var len1, len2 int
	len1 = len(num1)
	len2 = len(num2)
	strlen := len1
	if len2 > len1 {
		strlen = len2
	}
	s1 := ""
	sub := 0
	tmp1 := reverseString(num1)
	tmp2 := reverseString(num2)

	var a = 0
	var b = 0
	for i := 0; i < strlen; i++ {
		a = 0
		b = 0
		if i < len1 {
			a, _ = strconv.Atoi(string(tmp1[i]))
		}
		if i < len2 {
			b, _ = strconv.Atoi(string(tmp2[i]))
		}
		fmt.Println(a, b)
		t := a + b + sub
		if t > 9 {
			sub = 1
			s1 += strconv.Itoa(t - 10)
		} else {
			sub = 0
			s1 += strconv.Itoa(t)
		}
	}
	if sub > 0 {
		s1 += "1"
	}
	return reverseString(s1)
	// return s1
}
func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// @lc code=end

